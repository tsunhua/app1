# app1

将所有本地应用程序集成为一个应用。

技术栈: Vue3 + Gin

![Preivew](https://raw.githubusercontent.com/tsunhua/app1/main/priview.png)

## 启动

```bash
cd web
npm install
npm run build

cd ..
go run *.go
```

## 发布

1. 在 `app.prod.toml` 文件中配置您的应用程序。
2. 运行 `build.sh` 脚本。
3. 将 `build` 文件夹中的 `app1.linux.tar.gz` 归档文件复制并解压到您希望的位置。
4. 运行 `./app1`。
