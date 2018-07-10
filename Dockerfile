FROM golang:1.9-alpine3.7
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN apk add git
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN go build -o main .
CMD ["/app/main"]
