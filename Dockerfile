FROM registry.cn-hangzhou.aliyuncs.com/startops-base/golang:1.18.4

ADD Makefile Makefile
RUN apt update;apt install -y protobuf-compiler

RUN GOPROXY=https://goproxy.cn
RUN make init


