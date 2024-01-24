FROM golang:1.21-alpine as builder

ENV CGO_ENABLED=0
ENV GOPROXY=direct

RUN apk update
RUN apk add --no-cache git build-base

WORKDIR /src

COPY . .

RUN go build -o bin/ .

FROM alpine:3.19

WORKDIR /srv

COPY --from=builder /src/bin/mock-partner .

ENV HTTP_PORT ":8080"

EXPOSE 8080

ENTRYPOINT ["./mock-partner"]
