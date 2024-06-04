FROM golang:1.22.0
ENV GO111MODULE=off
WORKDIR /app

COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 go build GOOS=linux go build -o /testapp-logalerting

CMD ["/testapp-logalerting"]
