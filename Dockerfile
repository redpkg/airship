# build stage
FROM golang:1.13-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ARG GOPROXY
ENV GOPROXY=${GOPROXY}

WORKDIR $GOPATH/src/github.com/redpkg/airship

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags '-s -w' -o /app ./cmd/app/main.go

# final stage
FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /app .
COPY config.yml.example ./config.yml

ENTRYPOINT ["./app"]
