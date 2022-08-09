package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "myweb\t", log.LstdFlags)

func main() {
	logger.Println("Start server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	http.ListenAndServe(":8080", mux)
}

func greet(w http.ResponseWriter, r *http.Request) {
	logger.Println("greet")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
