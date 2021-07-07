FROM golang:1.14
WORKDIR /go/src/github.com/xdbfoundation/go

COPY . .
ENV GO111MODULE=on

# RUN go install github.com/xdbfoundation/go/tools/...
RUN go install github.com/xdbfoundation/go/services/...
