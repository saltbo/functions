FROM public.ecr.aws/lambda/provided:al2  AS build-env

ENV APP_HOME /app
WORKDIR $APP_HOME

# install compiler
RUN yum update -y \
    yum install -y yum-utils \
    yum-config-manager repo-pkgs -h \
    yum install -y golang g++ tesseract tesseract-langpack-eng

# todo 未安装成功，lambda的镜像到底是什么系统，什么版本。

ENV GOPROXY=https://goproxy.cn,direct
COPY go.* $APP_HOME/
RUN go mod download

COPY .. .
RUN go build -v -o server

# Runing environment
FROM public.ecr.aws/lambda/provided:al2

#RUN echo \
#    deb http://mirrors.aliyun.com/debian buster main \
#    deb http://mirrors.aliyun.com/debian buster-updates main \
#    deb http://mirrors.aliyun.com/debian-security buster/updates main \
#    > /etc/apt/sources.list

#    && yum install -y telnet libtesseract4 tesseract-ocr-eng

RUN yum update -y \
    yum install -y yum-utils \
    yum-config-manager -h \
    yum-config-manager --add-repo https://download.opensuse.org/repositories/home:/Alexander_Pozdnyakov/RHEL_7/ \
    yum install -y tesseract tesseract-langpack-eng

ENV APP_HOME /app
WORKDIR $APP_HOME

#COPY deployments/ $APP_HOME
COPY --from=build-env $APP_HOME/server /main

CMD ["/main"]