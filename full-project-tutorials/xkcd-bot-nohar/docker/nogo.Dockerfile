FROM golang AS builder
RUN mkdir -p $GOPATH/src/gitlab.com/neuware/xkcd-bot
WORKDIR $GOPATH/src/gitlab.com/neuware/xkcd-bot
ADD . .
RUN CGO_ENABLED=0 go build -o /bin/bot ./app

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /bin/bot /bin/bot
EXPOSE 2112
ENTRYPOINT ["/bin/bot"]
