package main

import (
	"app1/types"
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

var config types.Config
var siteMap map[string]types.Site
var userMap map[string]types.User

func init() {
	loadConfig()
}

func GetConfig() *types.Config {
	return &config
}

func GetSiteByPath(path string) (site types.Site, ok bool) {
	site, ok = siteMap[path]
	return
}

func GetSiteMap() map[string]types.Site{
	return siteMap
}

func GetUserByName(name string) (user types.User, ok bool) {
	user, ok = userMap[name]
	return
}

func loadConfig() {
	file, err := os.Open("app.toml")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Failed to decode config:", err)
		return
	}

	siteMap = make(map[string]types.Site)
	for _, v := range config.SiteGroups {
		for _, site := range v.Sites {
			siteMap[site.Path] = site
		}
	}

	userMap = make(map[string]types.User)
	for _, v := range config.Users {
		userMap[v.Name] = v
	}
}
