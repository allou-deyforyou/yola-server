FROM golang:1.18-alpine as builder

RUN apk add build-base
RUN mkdir /build
ADD ./ /build/
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/server /app/
COPY --from=builder /build/yola.db /app/
WORKDIR /app

CMD ["./server"]
