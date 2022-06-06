package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
)

var r map[string]interface{}

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

	log.Printf(strings.Repeat("-", 30))
	log.Printf("LIBS VERSION: %#+v\n", elasticsearch.Version)
	println()

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "PING")
	ping, _ := es.Ping()
	ping.Body.Close()

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "INFO")
	info, _ := es.Info()
	info.Body.Close()

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "CLUSTER")
	cluster, _ := es.Cluster.Health()
	cluster.Body.Close()

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "HEALTH")
	health, _ := es.Cat.Health(es.Cat.Health.WithV(true))
	log.Printf("health.String: \n%#+v\n", health.String())
	body, _ := ioutil.ReadAll(health.Body)
	log.Printf("health.Body: \n%#+v\n", string(body))
	health.Body.Close()
	println()

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "INDICES")
	indices, _ := es.Cat.Indices()
	log.Printf("indices.String: \n%#+v\n", indices.String())
	indices.Body.Close()
	println()

	/*
		log.Printf(strings.Repeat("-", 30))
		log.Printf("%#+v\n", "indicesStats")
		indicesStats, _ := es.Indices.Stats()
		indicesStats.Body.Close()
		println()
	*/

	log.Printf(strings.Repeat("-", 30))
	log.Printf("%#+v\n", "METRICS")
	metr, _ := es.Metrics()
	log.Printf("es.Metrics(): %#+v\n", metr)
}
