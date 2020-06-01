FROM golang:1.13  AS build-env

RUN echo \
    deb http://mirrors.aliyun.com/debian buster main \
    deb http://mirrors.aliyun.com/debian buster-updates main \
    deb http://mirrors.aliyun.com/debian-security buster/updates main \
    > /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y libtesseract-dev

ENV APP_HOME /app
WORKDIR $APP_HOME

ENV GOPROXY=https://goproxy.cn,direct
COPY go.* $APP_HOME/
RUN go mod download

COPY . .
RUN go build -v -o server cmd/server.go

# Runing environment
FROM debian:10

RUN echo \
    deb http://mirrors.aliyun.com/debian buster main \
    deb http://mirrors.aliyun.com/debian buster-updates main \
    deb http://mirrors.aliyun.com/debian-security buster/updates main \
    > /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y telnet libtesseract4 tesseract-ocr-eng

ENV APP_HOME /app
WORKDIR $APP_HOME

#COPY deployments/ $APP_HOME
COPY --from=build-env $APP_HOME/server $APP_HOME/bin/server

CMD ["bin/server"]