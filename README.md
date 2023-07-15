# app1

[简体中文](https://github.com/tsunhua/app1/blob/main/README-ZH.md) ｜ [繁體中文](https://github.com/tsunhua/app1/blob/main/README-ZHT.md)

Integrate ALL local application in ONE.

TECH: Vue3 + Gin

![Preivew](https://raw.githubusercontent.com/tsunhua/app1/main/priview.png)

## Startup

```bash
cd web
npm install
npm run build

cd ..
go run *.go
```

## Ship

1. config your applications in `app.prod.toml`.
2. run `build.sh`
3. copy and extract archive `app1.linux.tar.gz` in folder `build` to the location you want.
4. run `./app1`
