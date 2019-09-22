FROM golang:1.13.0-alpine3.10

COPY ./go.mod /linechatbot/go.mod
COPY ./app /linechatbot/app

RUN apk add --update --no-cache gcc git bind-dev && rm -rf /var/cache/apk/*
RUN go get -u github.com/oxequa/realize

WORKDIR /linechatbot/app

RUN CGO_ENABLED=0 go mod tidy
