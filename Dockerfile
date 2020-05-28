FROM golang:1.13  AS build-env

ENV APP_HOME /app
WORKDIR $APP_HOME

ENV GOPROXY=https://goproxy.cn,direct
COPY go.* $APP_HOME/
RUN go mod download

COPY . .
RUN go build -v -o server cmd/server.go

# Runing environment
FROM debian:9

RUN echo \
    deb http://mirrors.aliyun.com/debian/ stretch main non-free contrib\
    deb-src http://mirrors.aliyun.com/debian/ stretch main non-free contrib\
    deb http://mirrors.aliyun.com/debian-security stretch/updates main\
    deb-src http://mirrors.aliyun.com/debian-security stretch/updates main\
    deb http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib\
    deb-src http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib\
    deb http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib\
    deb-src http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib\
    > /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y telnet

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY deployments/ $APP_HOME
COPY --from=build-env $APP_HOME/server $APP_HOME/bin/server

CMD ["bin/server"]