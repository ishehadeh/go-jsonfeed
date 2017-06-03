package main

import (
	"log"
	"net/http"

	"github.com/IanS5/go-jsonfeed"
)

func main() {
	feed, err := jsonfeed.OpenFeed("https://jsonfeed.org/feed.json")
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", feed))
}
