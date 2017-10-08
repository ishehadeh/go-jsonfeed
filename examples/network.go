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

	//create a new text item with a single attachment
	helloItem := feed.PublishText("my-unique-id", "Hello World!", "Hello from https://github.com/IanS5/go-jsonfeed/examples/network.go!",
		jsonfeed.NewImage("cat", "https://static.pexels.com/photos/126407/pexels-photo-126407.jpeg"))

	//the item returned from feed.PublishXXX functions can be modified afterwards
	helloItem.Author = jsonfeed.Author{
		Name: "Ian Shehadeh",
	}

	log.Fatal(http.ListenAndServe(":8080", feed))
}
