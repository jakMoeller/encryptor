FROM docker.io/alpine:3.15 AS build
RUN apk add --update build-base ruby make

ADD . /tmp/build