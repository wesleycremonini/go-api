FROM golang:alpine as build

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/app

COPY ../go.mod /go/src/app
COPY ../go.sum /go/src/app
RUN go mod download

COPY ../ /go/src/app

RUN go build -v -ldflags "-s -w" -o /go/bin/app /go/src/app/cmd/api

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]

# docker build -t go_app -f docker/Dockerfile .
# docker run -p 80:80 go_app


# ----------- DEV -------------
# docker run --pull=missing -it --rm -w "/app" -v $(pwd):/app -p 80:80 --network="host" cosmtrek/air -c /app/docker/.air.toml