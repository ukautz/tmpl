FROM alpine:3.6

RUN apk update && apk add curl tar gzip libc6-compat

ARG GOVERSION=1.9
RUN curl -Lo /go.tar.gz https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz && \
    cd / && tar zxf go.tar.gz && rm go.tar.gz && \
    ln -s /go/bin/go /bin/go

ARG JQVERSION=1.5
RUN curl -Lo /bin/jq https://github.com/stedolan/jq/releases/download/jq-${JQVERSION}/jq-linux64 && \
    chmod +x /bin/jq

ADD https://github.com/lalyos/docker-upx/releases/download/v3.91/upx /bin/upx
RUN chmod +x /bin/upx

RUN mkdir -p /releases /opt/go /source /build/src/github.com/ukautz
ENV GOROOT=/go
ENV GOPATH=/build
ENV RELEASES=/releases
ENV SOURCE=/source
ENV USER_REPO=github.com/ukautz
ENV REPO=${USER_REPO}/tmpl
ENV ISBUILD=1

RUN mkdir -p $RELEASES ${GOPATH}/src/${USER_REPO} $SOURCE && \
    ln -s $SOURCE ${GOPATH}/src/${REPO}

VOLUME /source
VOLUME /releases


