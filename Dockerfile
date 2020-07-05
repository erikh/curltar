FROM golang:latest AS builder

RUN apt-get update -qq && apt-get install build-essential -y

WORKDIR /go/src/github.com/erikh/curltar
COPY . .

RUN go install -v ./...

FROM debian:latest

RUN apt-get update -qq && apt-get install ca-certificates -y && rm -rf /var/apt /var/cache
COPY --from=builder /go/bin/curltar /usr/bin/curltar

ENTRYPOINT [ "/usr/bin/curltar" ]
CMD []
