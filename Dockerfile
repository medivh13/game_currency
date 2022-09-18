FROM golang:1.17-alpine3.13 AS build-env

LABEL maintainer="jody almaida<jody.almaida@gmail.com>"

ENV APP_NAME=game-currency
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/medivh13
ENV TZ=Asia/Jakarta
ENV GIT_TERMINAL_PROMPT=0
ENV CGO_ENABLED=0

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    make \
    git \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add --no-cache coreutils

WORKDIR /src

RUN ls -ls

RUN mkdir -p /src/game-currency
COPY . /src/game-currency
WORKDIR /src/game-currency

RUN go mod tidy -compat=1.17

RUN go build

EXPOSE 8080

CMD "./game-currency"