FROM golang:1.17 as builder

WORKDIR /

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -ldflags '-s -w' -tags release -a -o server ./cmd/main.go

# multi stage build
FROM debian:buster-slim

ARG ENV=dev

COPY --from=builder /server /app/server
COPY --from=builder /env/${ENV}.env /app/env/${ENV}.env

CMD [ "./app/server", "-env=$ENV" ]