FROM golang:1.14-stretch AS builder

WORKDIR /go/src/github.com/xdbfoundation/go
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install github.com/xdbfoundation/go/services/frontier
RUN go install github.com/xdbfoundation/go/exp/services/captivecore

FROM ubuntu:18.04

ENV DIGITALBITS_CORE_VERSION 15.0.0-40
ENV DIGITALBITS_CORE_BINARY_PATH /usr/bin/digitalbits-core

ENV DEBIAN_FRONTEND=noninteractive
# ca-certificates are required to make tls connections
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl wget gnupg apt-utils
RUN wget -qO - https://apt.digitalbits.org/SDF.asc | APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=true apt-key add -
RUN echo "deb https://apt.digitalbits.org xenial stable" >/etc/apt/sources.list.d/SDF.list
RUN echo "deb https://apt.digitalbits.org xenial unstable" >/etc/apt/sources.list.d/SDF-unstable.list
RUN apt-get update && apt-get install -y digitalbits-core=${DIGITALBITS_CORE_VERSION}
RUN apt-get clean

COPY --from=builder /go/bin/frontier ./
COPY --from=builder /go/bin/captivecore ./

ENTRYPOINT ["./frontier"]
