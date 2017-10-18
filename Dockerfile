FROM golang:1.9-alpine

ADD . /go/src/app
WORKDIR /go/src/app


RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"] # ["app"]
