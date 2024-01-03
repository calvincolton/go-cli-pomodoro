FROM golang:1.21 AS builder
RUN mkdir /go-cli-pomodoro
COPY go-cli-pomodoro/ go-cli-pomodoro/
WORKDIR /go-cli-pomodoro
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers

FROM alpine:latest
RUN mkdir /app && adduser -h /app -D go-cli-pomodoro
WORKDIR /app
COPY --chown=go-cli-pomodoro --from=builder /go-cli-pomodoro .
CMD ["/app/go-cli-pomodoro"]