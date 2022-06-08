package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	fileName   string
	count      int
	batchSize  int        = 15
	batchLines [][]string = make([][]string, batchSize)
)

func init() {
	flag.StringVar(&fileName, "f", "", "csv file to index")
	flag.Parse()
}

// check https://github.com/elastic/go-elasticsearch/blob/v8.1.0/_examples/bulk/default.go
// for a more complete version
// this is basic
func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	cert, _ := ioutil.ReadFile("ca.crt")
	cfg := elasticsearch.Config{
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

	readFile(es, file)
}

func readFile(es *elasticsearch.Client, file io.ReadWriter) {
	// miss extract header and ignore it from batch
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				processBatch(es, batchLines[:count])
				break
			}
			log.Fatalln(err)
		}
		batchLines[count] = record
		count++
		if count == batchSize {
			processBatch(es, batchLines)
			count = 0
		}
	}
}

func processBatch(es *elasticsearch.Client, batch [][]string) {
	var buf bytes.Buffer

	end := []byte("")
	end = append(end, "\n"...)

	for _, v := range batch {
		// a completer
		buf.WriteString(`{ "index": { "_id":"` + v[0] + `"} }`)
		buf.Write(end)
		buf.WriteString(`{`)
		buf.WriteString(`"a":"`)
		buf.WriteString(v[0])
		buf.WriteString(`",`)
		buf.WriteString(`"b":"`)
		buf.WriteString(v[1])
		buf.WriteString(`",`)
		buf.WriteString(`"c":"`)
		buf.WriteString(v[2])
		buf.WriteString(`"}`)
		buf.Write(end)
	}

	res, err := es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex("test_bulk"))
	if err != nil {
		log.Fatalln("Error indexing batch: ", batch)
	}

	var raw map[string]interface{}
	json.NewDecoder(res.Body).Decode(&raw)
	if res.IsError() {
		log.Printf("Error: [%d] %s: %s",
			res.StatusCode,
			raw["error"].(map[string]interface{})["type"],
			raw["error"].(map[string]interface{})["reason"],
		)
	}
	log.Println("First document of each batch: ", raw["items"].([]interface{})[0])
	res.Body.Close()
}
