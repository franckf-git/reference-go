package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    serverMux5050 := http.NewServeMux()
    serverMux5050.HandleFunc("/", MainHandler)

    serverMux8080 := http.NewServeMux()
    serverMux8080.HandleFunc("/", MainHandler)

    fmt.Println("Listening on port 5050...")
    fmt.Println("Listening on port 8080...")

    go http.ListenAndServe(":5050", serverMux5050)
    http.ListenAndServe(":8080", serverMux8080)

}

