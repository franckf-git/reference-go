package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
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

	// curl -v --cacert ca.crt https://elastic:3lastic@localhost:9200
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error reading response: %s", err)
	}
	defer res.Body.Close()

	var r map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		log.Fatalf("Error decoding response: %s", err)
	}
	log.Printf("version: %#+v\n", r["version"])
	log.Printf("version-number: %#+v\n", r["version"].(map[string]interface{})["number"])

	log.Println(res)
}
