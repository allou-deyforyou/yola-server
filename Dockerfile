FROM golang:1.18-alpine

WORKDIR /src
COPY . .
RUN apk add build-base
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

EXPOSE 4000

CMD ["/src/server"]
