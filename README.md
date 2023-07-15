# app1

[简体中文](https://github.com/tsunhua/app1/blob/main/README-ZH.md) ｜ [繁體中文](https://github.com/tsunhua/app1/blob/main/README-ZHT.md)

Integrate ALL local applications in ONE.

TECH: Vue3 + Gin

![Preivew](https://raw.githubusercontent.com/tsunhua/app1/main/preview.png)

## Startup

```bash
cd web
npm install
npm run build

cd ..
go run *.go
```

Then open `http://127.0.0.1:8888`, enjoy app1.

## Ship

1. config your applications in `app.prod.toml`.
2. run `build.sh`
3. copy and extract archive `app1.linux.tar.gz` in folder `build` to the location you want.
4. run `./app1`
