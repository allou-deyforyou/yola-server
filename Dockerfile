FROM golang:1.18.3-alpine as build

RUN apk add build-base

RUN mkdir /build
ADD ./ /build/

WORKDIR /build

RUN go mod download
RUN GOPROXY=https://goproxy.io,direct go build -o server

FROM chromedp/headless-shell:latest

RUN apt-get update; apt-get upgrade; apt install -y dumb-init
ENTRYPOINT ["dumb-init", "--"]

COPY --from=build /build/server /app/
COPY --from=build /build/yola.db /app/

WORKDIR /app

CMD ["./server"]
