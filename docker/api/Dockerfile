FROM golang:1.15.3-alpine3.12
# Install basic packages
RUN apk update && \
    apk add build-base curl git make bash
RUN go get -u github.com/go-bindata/go-bindata/...
RUN go get -u github.com/oxequa/realize
WORKDIR /go/src/api