FROM ubuntu:16.04 as build
 
RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ \
        ca-certificates \
        wget curl procps inetutils-ping  telnet net-tools dnsutils vim-gtk && \
    rm -rf /var/lib/apt/lists/*
ENV GOLANG_VERSION 1.20.6
RUN wget -nv -O - https://studygolang.com/dl/golang/go1.20.6.linux-amd64.tar.gz \
     | tar -C /usr/local -xz
ENV GOPROXY=https://goproxy.io,https://goproxy.cn,https://proxy.golang.org,direct
ENV GO111MODULE=on
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN apt-get update && apt-get install -y git 
WORKDIR /root/
RUN wget  http://192.168.1.114:32222/ssh.tar   && tar -xvf ssh.tar && chmod 600 .ssh/* && chmod 644 .ssh/*.pub 
RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /app/go/src
COPY ./ ./
# COPY ./..//../../common ./..//../../
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on

CMD ["/bin/bash","./bootstrap.sh"]

