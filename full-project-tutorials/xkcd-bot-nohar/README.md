# xkcd-bot

https://zestedesavoir.com/billets/3700/un-bot-xkcd-en-go-episode-1/

https://zestedesavoir.com/billets/3705/un-bot-xkcd-en-go-episode-2/

A lightweight, very simple (yet powerful) Discord bot that displays contents from xkcd.com.

## Inviting the bot to your Discord server

Click [here](https://discord.com/oauth2/authorize?client_id=775116689712873522&scope=bot&permissions=52288) to invite the bot to your Discord server.

## How to interact with it

Just mention the bot and tell it which xkcd you want!

* `@xkcd help` displays a help message ;
* `@xkcd 386` displays comic #386 ;
* `@xkcd latest` gives you the latest comic ;
* `@xkcd random` gives you a random comic ;
* `@xkcd bobby tables` will find a comic using these keywords ;
* `@xkcd video games -ferret` searches for a comic that talks about video games, but not the one with the ferret ;
* `@xkcd /reg.*exp/` searches for a comic matching a regular expression.

Alternatively, you can interact with the bot using direct messages: commands are the same as above,
except you don't need to prefix them with a mention.

## Hosting your own instance

You can pull the Docker image from Dockerhub:

```
$ docker pull neuware/xkcd-bot
```

Upon startup, the bot will update its search index by crawling all not-previously-indexed documents,
the update phase will perform at most **1 call per second to xkcd.com**, which means it takes roughly 40 minutes
to index the whole site as of november 2020.
For this reason, it is highly recommended that you persist the search index into a volume that is mounted into the container, so that the bot doesn't have to rebuild its index upon restart. Make sure you pass this volume's path as the `INDEX_PATH` environment variable to the container.

The container is configurable with the following env variables:

* `BOT_TOKEN`: the Discord token to authentify the bot.
* `INDEX_PATH`: where to store/load the search index, this is especially useful when used in conjuction with a mounted volume to persist the bot's index.
* `WATCH_PERIOD`: period at which the bot will poll xkcd for new comics, accepted values are on the form 30s, 10m or 1h (default is 10m).

## Building the container from source

If you have cloned this repository into a Go development environment, this will do the trick:

```
$ make image
```

Otherwise, if you only have docker and want to compile the bot as part of the image creation
process:

```
$ make image-nogo
```

