#!/bin/sh
output="./tmp/dist/sbin"
commandName="web-framework"
entryFile="main.go" # 多个文件需要用空格分开

# 启动开发服务
dev() {
    fresh -c ./fresh_runner.conf
}

# 生成文档
initDoc(){
    swag init
}

doCompress() {
  if command -v upx >/dev/null 2>&1; then
      cd ${output}
      upx ${commandName}
      echo '已使用upx对可执行文件进行了压缩!!!'
    else
      echo '未对可执行文件进行压缩!!!'
  fi
}

# 编译项目
buildLinux() {
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o ${output}/${commandName} ${entryFile}

    # 执行压缩
    doCompress
}

buildWindow() {
    # https://github.com/mattn/go-sqlite3/issues/303
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o ${output}/${commandName} ${entryFile}

    # 执行压缩
    doCompress
}

buildMac() {
    CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o ${output}/${commandName} ${entryFile}

    # 执行压缩
    doCompress
}

case "$1" in
    dev)
        $1
        exit 0
        ;;
    initDoc)
        $1
        exit 0
        ;;
    buildLinux)
        $1
        exit 0
        ;;
    buildWindow)
        $1
        exit 0
        ;;
    buildMac)
        $1
        exit 0
        ;;
    *)
        echo $DIR
        echo 'Usage: {dev|buildLinux|buildWindow|buildMac}'
        ;;
esac
