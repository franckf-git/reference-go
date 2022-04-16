package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/gitlab", gitlab)
	http.HandleFunc("/linkedin", linkedin)
	http.ListenAndServe(":8888", nil)
}

func gitlab(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Body, r.Header, r.URL.Query(), r.Header.Get("X-REAL-IP"), r.Header.Get("X-FORWARDED-FOR"), r.RemoteAddr)
	http.Redirect(w, r, "https://gitlab.com/users/franckf/projects", http.StatusPermanentRedirect)
}
func linkedin(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Body, r.Header, r.URL.Query(), r.Header.Get("X-REAL-IP"), r.Header.Get("X-FORWARDED-FOR"), r.RemoteAddr)
	http.Redirect(w, r, "https://fr.linkedin.com/", http.StatusPermanentRedirect)
}
