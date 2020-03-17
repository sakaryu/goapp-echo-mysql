# Build Container
FROM golang:1.13-alpine3.10
# Build
RUN apk add --no-cache git
RUN go get github.com/oxequa/realize
RUN go get golang.org/x/lint/golint

WORKDIR /go/src
COPY . .
