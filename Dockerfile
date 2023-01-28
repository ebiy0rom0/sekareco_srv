FROM golang:1.18 as builder

WORKDIR /

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -ldflags '-s -w' -tags ${TAGS} -a -o serverd ./cmd/main.go

# multi stage build
FROM debian:buster-slim

ARG ENV=dev
ENV STAGE $ENV

COPY --from=builder /serverd /app/serverd

CMD ./app/serverd
