package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page
/**
 * The Body element is a []byte rather than string because that is the type expected by the io libraries
 */
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// The Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the request path. This is because the path will invariably begin with "/view/", which is not part of the page's title.
	title := r.URL.Path[len("/view/"):]
	// Again, note the use of _ to ignore the error return value from loadPage. This is done here for simplicity and generally considered bad practice.
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
