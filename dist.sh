#!/bin/bash
# Builds and packages the binaries for Darwin/amd64 and Linux/amd64

version=$(cat watchgopher.go | grep "const VERSION" | awk '{print $NF}' | sed 's/\"//g')
arch=$(go env GOARCH)

go test || (echo 'Tests failed. Stopping...' && exit 1)

for os in linux darwin; do
  target="watchgopher-$version-$os-$arch"
  echo "Building $target ..."

  mkdir -p dist/$target
  GOOS=$os go build -o dist/$target/watchgopher .
  cp README.md dist/$target

  cd dist
  tar czvf $target.tar.gz $target
  cd -
  rm -rf dist/$target
done
