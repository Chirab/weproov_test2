FROM golang:1.18-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY .env /app

RUN go build -o weproov cmd/weproov/weproov.go

EXPOSE 8083

CMD ["/app/weproov"]