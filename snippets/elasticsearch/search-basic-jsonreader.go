package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8@v8.2.0/esapi#Search.WithBody
// https://github.com/elastic/go-elasticsearch/blob/main/_examples/encoding/jsonreader.go

type requestQuery struct {
	Query query `json:"query"`
}
type query struct {
	QueryString queryString `json:"query_string"`
}
type queryString struct {
	Query string `json:"query"`
}

type errorResponse struct {
	Error errorF `json:"error"`
}
type errorF struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

type respQuery struct {
	Hits hits    `json:"hits"`
	Took float64 `json:"took"`
}
type hits struct {
	Total total `json:"total"`
	Hits  []hit `json:"hits"`
}
type total struct {
	Value float64 `json:"value"`
}
type hit struct {
	Index  string `json:"_index"`
	Id     string `json:"_id"`
	Source source `json:"_source"`
}
type source struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func main() {
	cert, _ := ioutil.ReadFile("ca.crt")
	cfg := elasticsearch.Config{
		Logger:            &elastictransport.CurlLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
		Addresses:         []string{"https://localhost:9200"},
		Username:          "elastic",
		Password:          "3lastic",
		EnableDebugLogger: true,
		EnableMetrics:     true,
		CACert:            cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var query = requestQuery{
		Query: query{
			QueryString: queryString{
				Query: "10*",
			},
		},
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test_indexer_nobatch"),
		es.Search.WithBody(esutil.NewJSONReader(&query)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e = errorResponse{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e.Error.Type,
				e.Error.Reason,
			)
		}
	}

	var r = respQuery{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r.Hits.Total.Value),
		int(r.Took),
	)
	for _, hit := range r.Hits.Hits {
		log.Printf("* ID=%s, %#v", hit.Id, hit.Source)
	}

}
