FROM golang:1.22.0
WORKDIR /app

COPY *.go ./

RUN GOOS=linux go build -o /testapp-logalerting

CMD ["/testapp-logalerting"]
