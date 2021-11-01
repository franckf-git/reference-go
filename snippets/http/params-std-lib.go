package main

import (
	"log"
	"net/http"
	"strings"
)

// curl 'localhost:8080/queryurl?a=red&b=blue'
func queryurl(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	for k, v := range values {
		log.Printf("k => v %#+v %v\n", k, v)
	}

	queryA := r.URL.Query().Get("a")
	queryB := r.URL.Query().Get("b")
	log.Printf("query: %#+v\n", queryA)
	log.Printf("query: %#+v\n", queryB)
}

// curl 'localhost:8080/reqform?a=red&b=blue'
func reqform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("query: %#+v\n", r.Form)
	log.Printf("queryA: %v\n", r.Form["a"])
	log.Printf("queryB: %v\n", r.Form["b"])
}

// curl localhost:8080/urlpath/3
func urlpath(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Path
	log.Printf("query: %#+v\n", query)
	id := strings.TrimPrefix(query, "/urlpath/")
	log.Printf("id: %#+v\n", id)
}

func urlpathRoot(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Path
	log.Printf("query: %#+v\n", query)
}

func main() {
	http.HandleFunc("/queryurl", queryurl)
	http.HandleFunc("/reqform", reqform)
	http.HandleFunc("/urlpath", urlpathRoot)
	http.HandleFunc("/urlpath/", urlpath)
	http.ListenAndServe(":8080", nil)
}
