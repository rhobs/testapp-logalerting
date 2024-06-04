FROM golang:1.22.0
ENV GO111MODULE=off
WORKDIR /app

COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD ["/testapp-logalerting"]
