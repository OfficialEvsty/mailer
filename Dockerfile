FROM golang:alpine AS builder

WORKDIR /mailer

COPY  go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd/
COPY config ./config/
COPY internal ./internal/
COPY domain ./domain/

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /mailer

COPY --from=builder /mailer/main /mailer/main

COPY config/local.yaml ./config/local.yaml

LABEL version="v0.1.0" author="evsty" desc="api mail service"

ENTRYPOINT ["/mailer/main"]