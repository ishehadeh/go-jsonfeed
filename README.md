Go JSON Feed
-----------------
Basic implementation of the [JSON Feed](https://jsonfeed.org/) specs for go.

Usage
------------
First install the package with `go get github.com/IanS5/go-jsonfeed`.

To get a feed use `feed,err := jsonfeed.OpenFeed(feed_url)`.

Feeds implement the `http.Handler` interface, so to serve a feed just use `http.Handle('/my/feed',feed)`.

Feeds can also be used with the `io.Writer` and `io.Reader` interfaces with their `Write(io.Writer)` and `Read(io.Reader)` methods.

All names in the JSONFeed struct have the same words as the json fields, but in CammelCase with the First letter capitalized.

Example
---------
```go
package main

import (
    "net/http"
    "github.com/IanS5/go-jsonfeed"
)
func main () {
    feed,err := jsonfeed.OpenFeed("https://jsonfeed.org/feed.json")
    if err != nil {
        panic(err)
    }
    log.Fatal(http.ListenAndServe(":8080",feed));
}
```