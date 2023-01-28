FROM golang:1.18 as builder

WORKDIR /

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -ldflags '-s -w' -tags release -a -o server ./cmd/main.go

# multi stage build
FROM debian:buster-slim

ARG ENV=dev
ENV STAGE $ENV

COPY --from=builder /server /app/server
COPY --from=builder /env/${ENV}.env /env/${ENV}.env

CMD ./app/server -stage=${STAGE}
