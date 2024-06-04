FROM golang:1.22.0
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /testapp-logalerting

CMD ["/testapp-logalerting"]
