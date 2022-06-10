package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
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

// check https://github.com/elastic/go-elasticsearch/blob/v8.1.0/_examples/bulk/indexer.go
// for a more complete version
// this is basic
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

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  "test_indexer_nobatch",
		Client: es,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	readFile(bi, file)

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	biStats := bi.Stats()
	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%v] documents with [%v]",
			biStats.NumFlushed,
			biStats.NumFailed,
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%v] documents",
			biStats.NumFlushed,
		)
	}
	dur := time.Since(start).Truncate(time.Millisecond)
	log.Printf("time.Since(start): %v\n", dur) // 667ms
}

func readFile(es esutil.BulkIndexer, file io.ReadWriter) {
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

			err = es.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					Action:     "index",
					DocumentID: "batch-0-numInBatch-" + strconv.Itoa(count),
					Body:       bytes.NewReader(buf.Bytes()),
					OnSuccess: func(
						ctx context.Context,
						item esutil.BulkIndexerItem,
						res esutil.BulkIndexerResponseItem,
					) {
						log.Printf("[%d] %s test/%s", res.Status, res.Result, item.DocumentID)
					},
					OnFailure: func(
						ctx context.Context,
						item esutil.BulkIndexerItem,
						res esutil.BulkIndexerResponseItem,
						err error) {
						if err != nil {
							log.Printf("ERROR: %s", err)
						} else {
							log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
						}
					},
				},
			)
			if err != nil {
				log.Fatalf("Unexpected error: %s", err)
			}
		}
		count++
	}
}
