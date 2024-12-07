#源镜像
FROM golang:latest
#FROM golang:1.17 as build
#作者
MAINTAINER Razil "test2022@163.com"
## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /www/webapp
## 设置工作目录
WORKDIR /www/webapp
## 把当前（宿主机上）目录下的文件都复制到docker上刚创建的目录下
COPY . /www/webapp
#将服务器的go工程代码加入到docker容器中
#ADD . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
RUN go build main.go
#暴露端口
EXPOSE 8080

RUN chmod +x main
ENTRYPOINT ["./main"]

## 启动docker需要执行的文件
#CMD go run main.go
#最终运行docker的命令
#ENTRYPOINT  ["./mygohttp"]