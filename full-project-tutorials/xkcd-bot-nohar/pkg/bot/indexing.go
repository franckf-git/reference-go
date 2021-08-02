package bot

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/neuware/xkcd-bot/pkg/index"
	"gitlab.com/neuware/xkcd-bot/pkg/xkcd"
)

func startIndexing(path string) {
	index.SetPath(path)
	xkcd.OnNewComic(index.AddComic)

	// start updating the index in the background
	go updateIndex()
}

func updateIndex() {
	log.Println("updating search index")
	latest := xkcd.LatestNum()
	for num := 1; num <= latest; num++ {
		if index.HasComic(num) {
			continue
		}
		c, err := xkcd.Get(fmt.Sprint(num))
		if err != nil {
			log.Printf("couldn't get #%d: %s", num, err)
			continue
		}
		index.AddComic(c)
		time.Sleep(time.Second)
	}
	log.Printf("search index up to date")
}
