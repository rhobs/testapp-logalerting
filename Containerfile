FROM golang:1.22.0
ENV GO111MODULE=off
WORKDIR /app

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /testapp-logalerting

CMD ["/testapp-logalerting"]
