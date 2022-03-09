FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=https://goproxy.cn,direct
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD .. .
RUN GOARCH=amd64 go build -o /main

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY --from=build /main /main
ENTRYPOINT [ "/main" ]