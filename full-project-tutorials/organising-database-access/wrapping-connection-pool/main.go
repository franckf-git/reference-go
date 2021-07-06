package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore.alexedwards.net/models"

	_ "github.com/lib/pq"
)

// This time make models.BookModel the dependency in Env.
type Env struct {
	books models.BookModel
}

func main() {
    // Initialise the connection pool as normal.
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

    // Initalise Env with a models.BookModel instance (which in turn wraps
    // the connection pool).
	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
    // Execute the SQL query by calling the All() method.
	bks, err := env.books.All()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
