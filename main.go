package main

import (
	"log"
	"manga-imdb/internal/site"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	site.PageRoutes(mux)
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
