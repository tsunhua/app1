// main.go
package main

import (
	"app1/log"
	"app1/types"
	"fmt"
	"net/http"
	"time"

	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置CORS中间件，允许跨域访问
	r.Use(cors.Default())

	// 登录接口
	api := r.Group("/api")
	api.POST("/login", loginHandler)

	// 站点列表接口
	api.GET("/sites", sitesHandler)

	// 代理请求接口
	for sitePath, site := range GetSiteMap() {
		func(sitePath string, site types.Site) {
			r.Any(fmt.Sprintf("/%s/*proxyPath", sitePath), authMiddleware(), func(ctx *gin.Context) {
				handleProxySite(ctx.Request, ctx.Writer, site)
			})
		}(sitePath, site)
	}
	r.NoRoute(func(c *gin.Context) {
		referer := c.Request.Header.Get("Referer")
		if referer != "" {
			if url, err := url.Parse(referer); err == nil {
				if sitePath := parseSitePath(url.Path); sitePath != "" {
					if site, ok := GetSiteByPath(sitePath); ok {
						handleProxySite(c.Request, c.Writer, site)
						return
					}
				}
			}
		}

		fs := http.FileServer(http.Dir("./web/dist"))
		writer := httptest.NewRecorder()
		req := c.Request.Clone(c)
		fs.ServeHTTP(writer, req)
		if writer.Code == http.StatusOK {
			c.Data(http.StatusOK, writer.HeaderMap["Content-Type"][0], writer.Body.Bytes())
			return
		}

		sitePath := parseSitePath(c.Request.URL.Path)
		// FOR SD
		if sitePath == "run" || sitePath == "queue" {
			handleProxySite(req, c.Writer, GetSiteMap()["draw"])
			return
		}
		for _, site := range GetSiteMap() {
			req := c.Request.Clone(c)
			writer := httptest.NewRecorder()
			handleProxySite(req, writer, site)
			if writer.Code == http.StatusOK {
				c.Data(http.StatusOK, writer.HeaderMap["Content-Type"][0], writer.Body.Bytes())
				return
			}
		}
	})
	r.Run(fmt.Sprintf("%s:%d", GetConfig().IP, GetConfig().Port))
}

func sitesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, GetConfig().SiteGroups)
}

func handleProxySite(req *http.Request, writer http.ResponseWriter, site types.Site) {
	var remote *url.URL
	var err error
	path := strings.TrimPrefix(req.URL.Path,"/" + site.Path)
	combinedUrl := fmt.Sprintf("http://%s:%d%s", GetConfig().IP, site.Port, path)
	remote, err = url.Parse(combinedUrl)
	if err != nil {
		log.Warn("parse url failed", "url", combinedUrl, "err", err.Error())
		return
	}
	// 目标站点的URL
	log.Debug("remote url: " + remote.String())
	// 创建反向代理器
	proxy := httputil.NewSingleHostReverseProxy(remote)
	// 修改请求头的 Host 字段，以便正确转发请求
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.Host = remote.Host
		req.URL.Path = path
	}
	// 启动服务器并进行代理转发
	proxy.ServeHTTP(writer, req)
}

func parseSitePath(proxyPath string) (sitePath string) {
	parts := strings.SplitN(proxyPath[1:], "/", 2)
	if len(parts) == 1 {
		sitePath = proxyPath[1:]
	} else {
		sitePath = parts[0]
	}
	return
}

func loginHandler(c *gin.Context) {
	var user types.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}
	inner, ok := GetUserByName(user.Name)
	if !ok || user.Password != inner.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}
	// 生成 JWT token
	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成 JWT"})
		return
	}

	// 返回 token
	c.JSON(http.StatusOK, gin.H{"app1_token": token})
}

// JWT 中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("app1_token")

		// 检查 token 是否存在
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供有效的认证凭证"})
			c.Abort()
			return
		}

		// 解析 token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 检查签名方法是否匹配
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("无效的签名方法")
			}
			return []byte(GetConfig().Secret), nil
		})

		// 检查解析结果是否有效
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的认证凭证"})
			c.Abort()
			return
		}

		// 将用户信息保存在上下文中，以便后续处理
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的认证凭证"})
			c.Abort()
			return
		}

		userID := claims["sub"].(string)
		c.Set("userID", userID)

		c.Next()
	}
}

// 生成 JWT token
func generateJWT(user types.User) (string, error) {
	// 设置 token 过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.Name,
		ExpiresAt: expirationTime.Unix(),
	})

	// 签名 token 并返回字符串形式
	tokenString, err := token.SignedString([]byte(GetConfig().Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
