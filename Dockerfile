FROM golang:1.18-alpine as builder

RUN apk add build-base
RUN mkdir /build
ADD ./ /build/
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM chromedp/headless-shell:latest

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/server /app/
COPY --from=builder /build/yola.db /app/
WORKDIR /app

CMD ["./server"]
