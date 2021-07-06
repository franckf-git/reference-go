package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "bookstore.alexedwards.net/models"

    _ "github.com/lib/pq"
)

func main() {
    var err error

    // Initalize the sql.DB connection pool and assign it to the models.DB 
    // global variable.
    models.DB, err = sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/books", booksIndex)
    http.ListenAndServe(":3000", nil)
}

// booksIndex sends a HTTP response listing all books.
func booksIndex(w http.ResponseWriter, r *http.Request) {
    bks, err := models.AllBooks()
    if err != nil {
        log.Println(err)
        http.Error(w, http.StatusText(500), 500)
        return
    }

    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}
