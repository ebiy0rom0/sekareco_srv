FROM golang:1.18 as builder
ARG tags=0

WORKDIR /

COPY go.* ./
RUN go mod download

COPY . ./

RUN apt-get install -y make

RUN make build RELEASE=${tags}

# multi stage build
FROM debian:buster-slim

COPY --from=builder /bin/serverd /app/serverd
COPY docs/db/ ./docs/db

EXPOSE 8000
CMD ./app/serverd
