FROM golang:1.22.0
WORKDIR /app

COPY *.go ./
COPY . ./

RUN go build GOOS=linux go build -o /testapp-logalerting

CMD ["/testapp-logalerting"]
