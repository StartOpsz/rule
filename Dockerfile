FROM registry.cn-hangzhou.aliyuncs.com/startops-base/golang:1.20

ADD Makefile Makefile
RUN apt update;apt install -y protobuf-compiler

ENV GOPROXY=https://goproxy.cn
RUN make init


