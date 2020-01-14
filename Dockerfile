FROM golang:1.13-alpine3.10 AS build

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod init

RUN CGO_ENABLED=0 go build -o go-app main.go

FROM busybox

COPY --from=build /go/src/app/go-app /usr/local/bin/go-app

ENTRYPOINT ["/usr/local/bin/go-app"]