FROM alpine:latest

RUN mkdir /app && adduser -h /app -D go-cli-pomodoro
WORKDIR /app
COPY --chown=go-cli-pomodoro /go-cli-pomodoro .
CMD ["/app/go-cli-pomodoro"]
