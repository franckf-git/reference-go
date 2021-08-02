all: bot

.PHONY: bot

bot:
	CGO_ENABLED=0 go build -o bin/bot ./app

image: bot
	docker build -f docker/Dockerfile -t neuware/xkcd-bot .

image-nogo:
	docker build -f docker/nogo.Dockerfile -t neuware/xkcd-bot .

publish:
	docker push neuware/xkcd-bot
