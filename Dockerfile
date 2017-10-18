FROM golang:1.9-alpine

RUN apk add --no-cache git

ADD . /go/src/app
WORKDIR /go/src/app


RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"] # ["app"]
