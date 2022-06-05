package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

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

	for i, v := range []string{"test 1", "test 2"} {
		data, _ := json.Marshal(struct{ Title string }{Title: v}) // keep field in Uppercase
		req := esapi.IndexRequest{
			Index:      "test_part_1",
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
	}

	for i, v := range []string{"test 3", "test 4"} {
		var b strings.Builder
		b.WriteString(`{"title":"`)
		b.WriteString(v)
		b.WriteString(`"}`)
		req := esapi.IndexRequest{
			Index:      "test_part_2",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(b.String()),
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
	}

	type testDocument struct {
		Title string `json:"title"`
	}
	for i, v := range []string{"test 5", "test 6"} {
		data := testDocument{Title: v}
		req := esapi.IndexRequest{
			Index:      "test_part_3",
			DocumentID: strconv.Itoa(i + 1),
			Body:       esutil.NewJSONReader(data),
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
	}
}
