# app1

整合所有本地應用程式為一個應用程式。

技術: Vue3 + Gin

![Preivew](https://raw.githubusercontent.com/tsunhua/app1/main/preview.png)

## 啟動

```bash
cd web
npm install
npm run build

cd ..
go run *.go
```

## 發佈

1. 在 `app.prod.toml` 中設定您的應用程式。
2. 執行 `build.sh` 腳本。
3. 將 `build` 資料夾中的 `app1.linux.tar.gz` 壓縮檔複製並解壓縮到您希望的位置。
4. 執行 `./app1`。
