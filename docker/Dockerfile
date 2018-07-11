FROM ubuntu:xenial-20180525

COPY sources.list /etc/apt/sources.list
COPY go1.10.3.linux-amd64.tar.gz /go.tgz
COPY bash_aliases /root/.bash_aliases
COPY libs/usr/lib/ /usr/lib
COPY dstrad /root/
COPY trawd /root/

RUN \
 apt-get update && \
 apt-get -y install git; \
 tar -C /usr/local -xzvf go.tgz; \
 rm go.tgz; \
 export PATH="/usr/local/go/bin:$PATH"; \
 go version;

ENV GOLANG_VERSION 1.10.2
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH
