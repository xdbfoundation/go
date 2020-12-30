FROM ubuntu:18.04

MAINTAINER Bartek Nowotarski <bartek@digitalbits.org>

ENV DIGITALBITS_CORE_VERSION 15.0.0-40
# to remove tzdata interactive flow
ENV DEBIAN_FRONTEND=noninteractive

ADD dependencies /
RUN ["chmod", "+x", "dependencies"]
RUN /dependencies

ADD digitalbits-core.cfg /

ADD start /
RUN ["chmod", "+x", "start"]

ENTRYPOINT ["/start"]