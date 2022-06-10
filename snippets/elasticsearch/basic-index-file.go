package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	fileName string
	count    int
	header   []string
)

func init() {
	flag.StringVar(&fileName, "f", "", "csv file to index")
	flag.Parse()
}

func main() {
	start := time.Now().UTC()
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
	dur := time.Since(start)
	log.Printf("time.Since(start): %v\n", dur) // 21.986778965s
}

func readFile(es *elasticsearch.Client, file io.ReadWriter) {
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		if count == 0 {
			header = record
		} else { // ugly else but don't find another way to skip header - 'continue' don't make it with infinite loop
			var buf bytes.Buffer
			buf.WriteString(`{`)
			buf.WriteString(`"` + header[0] + `":"`)
			buf.WriteString(record[0])
			buf.WriteString(`",`)
			buf.WriteString(`"` + header[1] + `":"`)
			buf.WriteString(record[1])
			buf.WriteString(`",`)
			buf.WriteString(`"` + header[2] + `":"`)
			buf.WriteString(record[2])
			buf.WriteString(`"}`)

			req := esapi.IndexRequest{
				Index:      "test_basic",
				DocumentID: strconv.Itoa(count),
				Body:       strings.NewReader(buf.String()),
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
		count++
	}
}
