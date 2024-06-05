package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tars47/go-url-shortener-boltdb/store"
	"github.com/tars47/go-url-shortener-boltdb/urlshort"
)

func main() {

	db, err := store.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", &urlshort.Service{Store: db})

	fmt.Println("--- Starting the server on Port 4747 ---")
	log.Fatal(http.ListenAndServe(":4747", nil))
}
