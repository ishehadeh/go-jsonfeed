Go JSON Feed
-----------------
Basic implementation of the [JSON Feed](https://jsonfeed.org/) specs for go.

Usage
------------
First install the package with `go get github.com/IanS5/go-jsonfeed`.

To get a feed use `feed, err := jsonfeed.OpenFeed(feed_url)`.

Feeds implement the `http.Handler` interface, so to serve a feed just use `http.Handle('/my/feed/url', feed)`.

Feeds can write data with the `JSONFeed.Write(io.Writer)` method and read data with the `JSONFeed.Read(io.Reader)` method.

All names in the JSONFeed struct have the same words as the json fields, but in CammelCase with the First letter capitalized, URL and ID are always all uppercase. `home_page_url` = `HomePageURL`

Example
---------

A more complicated example can be found in `examples/network.go`

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