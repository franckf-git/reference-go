package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
)

// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8@v8.2.0/esapi#Get

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

	res, err := es.Get(
		"test_indexer_nobatch",
		"batch-0-numInBatch-103",
		es.Get.WithPretty(),
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

	var r = hit{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s]",
		res.Status(),
	)
	log.Printf("* ID=%s, %#v", r.Id, r.Source)
}
