FROM golang:1.18.3-alpine

RUN apk add chromium

RUN mkdir /build
ADD ./ /build/

WORKDIR /build

RUN go mod download
RUN go build -o server

CMD ["./server"]

# RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

# FROM chromedp/headless-shell:latest

# RUN apt update; apt upgrade; apt install dumb-init

# ENTRYPOINT ["dumb-init", "--"]

# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/server /app/
# COPY --from=builder /build/yola.db /app/
# WORKDIR /app

