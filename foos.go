package main

import (
	s "github.com/gabesullice/foos/storage"
	"log"
	"net/http"
)

var storage *s.Storage

func init() {
	storage = s.NewStorage()
}

func main() {
	r := NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
