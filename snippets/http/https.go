package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTPS network address")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Proudly served with Go and HTTPS!")
	})
	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}
	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServeTLS(*certFile, *keyFile)
	log.Fatal(err)
}

// ./bin/go run src/crypto/tls/generate_cert.go -ecdsa-curve P256 -host localhost:4000
// openssl x509 -in cert.pem -text
// curl -Lv --cacert cert.pem https://localhost:4000
