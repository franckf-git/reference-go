FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD ./bin/bot /bin/bot
EXPOSE 2112
ENTRYPOINT ["/bin/bot"]
