FROM golang:1.18.3-alpine

RUN apk add build-base chromium

RUN mkdir /build
ADD ./ /build/

WORKDIR /build

RUN go mod download
RUN GOPROXY=https://goproxy.io,direct go build -o server

CMD ["./server"]
