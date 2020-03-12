# Build Container
FROM golang:1.13-alpine3.10 AS build
# Build
RUN apk add --no-cache git
RUN go get github.com/oxequa/realize

WORKDIR /go/src
COPY . .
RUN go build -o go-app

# Runtime Container
FROM alpine:3.10
COPY --from=build /go/src/app/go-app /usr/local/bin/go-app
COPY --from=build /go/src/app/start_app.sh /usr/local/bin/start_app.sh
ENTRYPOINT ["/bin/sh", "/usr/local/bin/start_app.sh"]