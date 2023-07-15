# app1

Integrate ALL local application in ONE.

TECH: Vue3 + Gin

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
