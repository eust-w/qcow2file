#!/bin/bash

CGO_ENABLED_ORI=`go env CGO_ENABLED`
GOOS_ORI=`go env GOOS`
GOARCH_ORI=`go env GOARCH`

go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
go env -w GO111MODULE=on
go env -w  GOPROXY=https://goproxy.cn,direct
cd ../
go mod tidy
mkdir pkg
go build -ldflags '-w -s' -trimpath -gcflags '-l' -a -o ./pkg/qcow2file

chmod 777 ./pkg/qcow2file
go env -w CGO_ENABLED=$CGO_ENABLED_ORI
go env -w GOOS=$GOOS_ORI
go env -w GOARCH=$GOARCH_ORI
cd ./build || exit
echo "qcow2filez_linux_amd64 build success"
