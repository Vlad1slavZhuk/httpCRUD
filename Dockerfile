FROM golang:alpine AS builder
LABEL maintainer="Vladyslav Zhuchkov"
WORKDIR /source
COPY . /source/
EXPOSE 8081


