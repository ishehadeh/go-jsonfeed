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
	helloItem := feed.PublishText("my-unique-id", "Hello World!", "Hello from https://github.com/IanS5/go-jsonfeed!",
		jsonfeed.NewImage("cat", "https://static.pexels.com/photos/126407/pexels-photo-126407.jpeg"))

	helloItem.Author = jsonfeed.Author{
		Name: "Ian Shehadeh",
	}

	log.Fatal(http.ListenAndServe(":8080", feed))
}
