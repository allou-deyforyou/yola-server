FROM golang:1.18-alpine as builder

RUN apk add build-base
RUN mkdir /build
ADD ./ /build/
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine

RUN apk add build-base

RUN curl -SL https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /google-chrome-stable_current_amd64.deb && \
    dpkg -x /google-chrome-stable_current_amd64.deb google-chrome-stable && \
    mv /google-chrome-stable/usr/bin/* /usr/bin && \
    mv /google-chrome-stable/usr/share/* /usr/share && \
    mv /google-chrome-stable/etc/* /etc && \
    mv /google-chrome-stable/opt/* /opt && \
    rm -r /google-chrome-stable


RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/server /app/
COPY --from=builder /build/yola.db /app/
WORKDIR /app

CMD ["./server"]
