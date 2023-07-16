#!/bin/bash
workdir=$(pwd)
os=linux
target="${workdir}/build/${os}"

echo "[1] Building web"

mkdir -p "${target}/web"
cd web
npm run build
rm -rf "${target}/web/dist"
cp -rf dist "${target}/web"

echo "[2] Building app1"
cd "${workdir}"
cp app.prod.toml build/${os}/app.toml
GOOS=${os} GOARCH=amd64 go build -o build/${os}/app1

echo "[3] Tar ${os}"
cd build
mkdir app1
cp -rf ${os}/ app1
tar -czvf app1.${os}.tar.gz app1/*
rm -rf app1

echo "✓ Build completed!"
