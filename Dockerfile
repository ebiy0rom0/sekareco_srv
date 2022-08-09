FROM golang:1.17 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -ldflags '-s -w' -tags release -a -o server ./cmd/main.go

# multi stage build
FROM debian:buster-slim

COPY --from=builder /app/server /app/server

CMD [ "/app/server" ]