package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "bookstore.alexedwards.net/models"

    _ "github.com/lib/pq"
)

// Create some middleware which swaps out the existing request context
// with new context.Context value containing the connection pool.
func injectDB(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "db", db)

        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

func main() {
    db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Fatal(err)
    }

    // Wrap the booksIndex handler with the injectDB middleware,
    // passing in the new context.Context with the connection pool.
    http.Handle("/books", injectDB(db, booksIndex))
    http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
    // Pass the request context onto the database layer.
    bks, err := models.AllBooks(r.Context())
    if err != nil {
        log.Println(err)
        http.Error(w, http.StatusText(500), 500)
        return
    }

    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}
