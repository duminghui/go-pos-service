FROM ubuntu:xenial-20180417

COPY go1.10.2.linux-amd64.tar.gz /go.tgz
COPY miniupnpc-1.9.20140401.tar.gz /
COPY sources.list /etc/apt/sources.list
COPY bash_aliases /root/.bash_aliases
COPY dstrad /root/
COPY trawd /root/

RUN \
 apt-get update && \
 # apt-get -y install software-properties-common && \
 \
 apt-get -y install libboost-all-dev libdb++-dev make git bash-completion; \
 tar -xzvf miniupnpc-1.9.20140401.tar.gz ; \
 cd miniupnpc-1.9.20140401 ; \
 make ; \
 make install ; \
 cd / ; \
 rm -rf miniupnpc-1.9.20140401.tar.gz miniupnpc-1.9.20140401 ; \
 tar -C /usr/local -xzvf go.tgz; \
 rm go.tgz; \
 export PATH="/usr/local/go/bin:$PATH"; \
 go version; \
 \
 apt-get -y autoremove ; \
 apt-get clean ; \
 rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.10.2
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH
