FROM golang:1.18.5-alpine3.9 AS builder

RUN apk update && apk upgrade && \
    apk --update add git make

workdir /.

RUN go mod tidy
RUN go build -o mainrun

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata


workdir /.

STOPSIGNAL SIGINT
EXPOSE 8080
COPY --from=builder /mainrun .
CMD ["./mainrun"]
