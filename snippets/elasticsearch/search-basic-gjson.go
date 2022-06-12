package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
)

// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8@v8.2.0/esapi#Search.WithBody
// https://www.elastic.co/fr/blog/the-go-client-for-elasticsearch-working-with-data

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

	body := `{"query":{"query_string":{"query":"10*"}}}`

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test_indexer_nobatch"),
		es.Search.WithBody(strings.NewReader(body)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	var respBuffer bytes.Buffer
	respBuffer.ReadFrom(res.Body)

	if res.IsError() {
		values := gjson.GetManyBytes(respBuffer.Bytes(), "error.type", "error.reason")
		log.Fatalf("[%s] %s: %s",
			res.Status(),
			values[0].String(),
			values[1].String(),
		)
	}

	values := gjson.GetManyBytes(respBuffer.Bytes(), "hits.total.value", "took", "hits.hits")

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		values[0].Int(),
		values[1].Int(),
	)

	for _, hit := range values[2].Array() {
		source := hit.Map()["_source"].Map()
		log.Printf("hit: a: %s - b: %s - c: %s \n", source["a"].String(), source["b"].String(), source["c"].String())
	}

}
