FROM alpine

RUN apk add chromium

RUN mkdir /build
ADD ./build/* /build/
ADD ./yola.db  /build/

WORKDIR /build

CMD ["./yola"]

# RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

# FROM chromedp/headless-shell:latest

# RUN apt update; apt upgrade; apt install dumb-init

# ENTRYPOINT ["dumb-init", "--"]

# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/server /app/
# COPY --from=builder /build/yola.db /app/
# WORKDIR /app
