FROM golang:1.18.3-alpine AS builder

RUN apk add build-base

RUN mkdir /build
ADD ./ /build/

WORKDIR /build

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM chromedp/headless-shell:latest

RUN apt-get update; apt-get upgrade; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]

COPY --from=builder /build/server /app/
COPY --from=builder /build/yola.db /app/

WORKDIR /app

CMD ["./server"]
