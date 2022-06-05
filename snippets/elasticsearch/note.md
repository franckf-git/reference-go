- Empirically, JSON encoding and decoding has the biggest influence on the performance of the client

+ https://pkg.go.dev/github.com/elastic/go-elasticsearch/v7/esutil#NewJSONReader
```
type MyDocument struct {
  Title string `json:"title"`
}

doc := MyDocument{Title: "Test"}

res, _ := es.Index("test", esutil.NewJSONReader(&doc))
fmt.Println(res)
// [201 Created] {"_index":"test","_id":"wleUWXQBe...
```

- https://www.elastic.co/fr/blog/the-go-client-for-elasticsearch-working-with-data

+ https://pkg.go.dev/github.com/elastic/go-elasticsearch/v7@v7.8.0/estransport?tab=doc#CurlLogger

+ Set the EnableMetrics option to true, and use the Metrics() method to retrieve the information. The examples/instrumentation/expvar.go example shows an integration with the expvar package.

- tidwall/gjson package. It allows you to use the "dot notation" — familiar from the jq command line utility — to "pluck" the values from the response easily, as well as more efficiently:

```
var b bytes.Buffer

res, _ := es.Search(es.Search.WithTrackTotalHits(true))
b.ReadFrom(res.Body)

values := gjson.GetManyBytes(b.Bytes(), "hits.total.value", "took")
fmt.Printf(
  "[%s] %d hits; took: %dms\n",
  res.Status(),
  values[0].Int(),
  values[1].Int(),
)
// => [200 OK] 1 hits; took: 10ms
```
