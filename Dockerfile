FROM golang:alpine AS builder
LABEL maintainer="Vladyslav Zhuchkov"
WORKDIR /source
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o server /source/cmd/main.go

EXPOSE 8081

FROM alpine
WORKDIR /root
COPY --from=builder /source/server .
CMD [./server]

