FROM golang:1.14
WORKDIR /go/src/github.com/digitalbits/go

COPY . .
ENV GO111MODULE=on

# RUN go install github.com/digitalbits/go/tools/...
RUN go install github.com/digitalbits/go/services/...
