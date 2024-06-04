package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tars47/go-url-shortener-boltdb/store"
	"github.com/tars47/go-url-shortener-boltdb/urlshort"
)

func main() {

	store, err := store.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	us := urlshort.Service{
		Store: store,
	}

	http.HandleFunc("/", us.Handle)

	fmt.Println("Starting the server on :4747")
	log.Fatal(http.ListenAndServe(":4747", nil))
}
