package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var wg sync.WaitGroup

// concurrency on limited slice, so similarity with bulk import
func main() {
	cert, _ := ioutil.ReadFile("ca.crt")
	cfg := elasticsearch.Config{
		Logger:            &elastictransport.CurlLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
		Addresses:         []string{"https://localhost:9200"},
		Username:          "elastic",
		Password:          "3lastic",
		EnableDebugLogger: true,
		CACert:            cert,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	for i, v := range []string{"test conc 1", "test conc 2"} {
		wg.Add(1)

		go func(i int, v string) {
			defer wg.Done()

			data, _ := json.Marshal(struct{ Title string }{Title: v})
			req := esapi.IndexRequest{
				Index:      "test_part_concurrent",
				DocumentID: strconv.Itoa(i + 1),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error res indexing: %s", err)
			}
			defer res.Body.Close()
			if res.IsError() {
				log.Fatalf("Error indexing: %s", err)
			}
			var r map[string]interface{}
			err = json.NewDecoder(res.Body).Decode(&r)
			if err != nil {
				log.Fatalf("Error res decoding: %s", err)
			}
			log.Printf("response: %#+v \n result: %#+v", res.Status(), r["result"])
		}(i, v)

		wg.Wait()
	}

}
