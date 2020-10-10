FROM golang:alpine AS builder
LABEL maintainer="Vladyslav Zhuchkov"
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o server /app/cmd/main.go

FROM alpine
WORKDIR /app
COPY template /app/template
COPY --from=builder /app/server .
CMD ["./server"]
