package main

import (
	"log"
	"net/http"

	jsonfeed "github.com/IanS5/go-jsonfeed"
)

func main() {
	feed, err := jsonfeed.OpenFeed("https://jsonfeed.org/feed.json")
	if err != nil {
		panic(err)
	}
	feed.PublishText("my-unique-id", "Hello World!", "Hello from https://github.com/IanS5/go-jsonfeed!")
	log.Fatal(http.ListenAndServe(":8080", feed))
}
