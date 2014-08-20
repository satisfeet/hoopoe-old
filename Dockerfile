FROM golang:1.3.1

RUN mkdir -p /go/src/github.com/satisfeet/hoopoe

ADD . /go/src/github.com/satisfeet/hoopoe

WORKDIR /go/src/github.com/satisfeet/hoopoe

RUN go get ./...

EXPOSE 80

CMD ["go", "run", "main.go", "--host", ":80", "--mongo", "mongo/satisfeet"]
