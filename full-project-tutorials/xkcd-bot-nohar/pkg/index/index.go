package index

import (
	"fmt"
	"log"
	"sync"

	"github.com/blevesearch/bleve"
	"gitlab.com/neuware/xkcd-bot/pkg/xkcd"
)

var (
	path = "xkcd.index"
	idx  bleve.Index
	once sync.Once
)

// open the index, creating it if it doesn't exist
func open() {
	var err error
	if idx, err = bleve.Open(path); err == nil {
		return
	}
	log.Printf("creating index at %q", path)
	m := bleve.NewIndexMapping()
	idx, err = bleve.New(path, m)
	if err != nil {
		log.Fatalf("index creation failed: %s", err)
	}
}

// SetPath sets the path of the current index.
func SetPath(p string) {
	path = p
	if idx != nil {
		open()
	}
}

// HasComic returns true if a comic with given ID exists in the current index.
func HasComic(id int) bool {
	once.Do(open)
	doc, _ := idx.Document(fmt.Sprint(id))
	return doc != nil
}

// AddComic adds a comic into the current index.
func AddComic(c *xkcd.Comic) {
	once.Do(open)
	id := fmt.Sprint(c.Num)
	c.Transcript = c.Script()
	idx.Index(fmt.Sprint(c.Num), c)
	log.Printf("indexed comic #%s", id)
}

// A SearchResult is the association of a comic number and its score for the
// search.
type SearchResult struct {
	ComicNum string
	Score    float64
}

// Search performs a full-text search on the current index.
func Search(query string) []SearchResult {
	once.Do(open)
	results, err := idx.Search(bleve.NewSearchRequest(bleve.NewQueryStringQuery(query)))
	if err != nil {
		log.Printf("while searching for %q: %s", query, err)
		return []SearchResult{}
	}
	res := make([]SearchResult, len(results.Hits))
	for i, r := range results.Hits {
		res[i] = SearchResult{r.ID, r.Score}
	}
	return res
}
