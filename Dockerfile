FROM golang:alpine AS base

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED="0"

RUN apk update \
  && apk add --no-cache \
  ca-certificates \
  git \
  && update-ca-certificates


###

FROM base AS dev

WORKDIR /app

RUN go get -u github.com/cosmtrek/air \
  && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air", "-c", "./config/.air.toml"]

###

FROM base AS builder

WORKDIR /app
COPY . /app

RUN go build -o soccer-bet-api -a .

### Production
FROM alpine:latest

ENV PATH="${PATH}:/sbin"

RUN apk update \
  && apk add --no-cache \
  ca-certificates \
  curl \
  tzdata \
  && update-ca-certificates

# Copy executable
COPY --from=builder /app/soccer-bet-api /usr/local/bin/soccer-bet-api
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/soccer-bet-api"]