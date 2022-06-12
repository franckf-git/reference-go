package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8@v8.2.0/esapi#SQLQuery

type errorResponse struct {
	Error errorF `json:"error"`
}
type errorF struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

type respQuery struct {
	Columns []column   `json:"columns"`
	Rows    [][]string `json:"rows"`
}
type column struct {
	Name string `json:"name"`
	Type string `json:"type"`
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

	body := `{"query":"\nSELECT * FROM _all WHERE a LIKE '10%' \n"}`
	reqSQL := esapi.SQLQueryRequest{
		Body:   strings.NewReader(body),
		Format: "json",
		Pretty: true,
	}

	res, err := reqSQL.Do(context.Background(), es)
	// or
	// res, err := es.SQL.Query(strings.NewReader(body))

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
		"[%s] %d hits;",
		res.Status(),
		len(r.Rows),
	)
	for _, headers := range r.Columns {
		log.Printf("* , %#v", headers.Name)
	}
	for _, row := range r.Rows {
		log.Printf("* , %#v", row)
	}

}
