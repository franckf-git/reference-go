package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var mappingsFields = `
{
"mappings": {
"properties": {
"session":                   { "type": "integer" },
"contrat":                   { "type": "text" },
"geolocalisation":           { "type": "geo_point" }
              }
            }
}`

func main() {
	index := "mon-index-avec-mapping"

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

	createIndex(es, index, mappingsFields)
	createDataViewKibana(index)

	// add document to valid
	type testDocument struct {
		Session     int    `json:"session"`
		Contrat     string `json:"contrat"`
		Geolocation string `json:"geolocalisation"`
	}
	data := testDocument{
		Session:     2021,
		Contrat:     "public",
		Geolocation: "48.6551,6.15225",
	}
	req := esapi.IndexRequest{
		Index:   index,
		Body:    esutil.NewJSONReader(data),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error res indexing: %s", err)
	}
	defer res.Body.Close()
}

func createIndex(es *elasticsearch.Client, indexName string, mapping string) {
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("error: Indices.Exists: %s", err)
	}
	res.Body.Close()
	if res.StatusCode == 404 {
		res, err := es.Indices.Create(
			indexName,
			es.Indices.Create.WithBody(strings.NewReader(mapping)),
		)
		if err != nil {
			log.Fatalf("error: Indices.Create: %s", err)
		}
		if res.IsError() {
			log.Fatalf("error: Indices.Create: %s", res)
		}
	}
}

func createDataViewKibana(index string) {
	data := fmt.Sprintf(`{"data_view":{"title":"%v*"}}`, index)

	req, err := http.NewRequest("POST", "http://localhost:5601/api/data_views/data_view", strings.NewReader(data))
	if err != nil {
		log.Fatalf("error while requesting kibana: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("kbn-xsrf", "true")
	req.SetBasicAuth("elastic", "3lastic")
	client := &http.Client{}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("resp.Body: \n%#+v\n", string(body))
	if err != nil {
		log.Fatalf("error while requesting kibana: %s", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("error kibana bad response - %d: %s", resp.StatusCode, err)
	}
	resp.Body.Close()
}
