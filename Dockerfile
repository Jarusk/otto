FROM alpine:3.6

RUN apk add --no-cache bash curl openssl-dev g++ python2 cmake git file make

WORKDIR /code

RUN git clone https://github.com/rust-lang/rust.git

WORKDIR rust

RUN git checkout tags/1.20.0

#RUN ./x.py build

#RUN ./x.py install
