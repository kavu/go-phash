# go-phash #

**go-phash** is a simple [pHash](http://phash.org) wrapper library for the Go programming language.

You can view documentation on [godoc.org](http://godoc.org/github.com/kavu/go-phash "go-phash documentation").
.

## Example ##

First of all, install **go-phash** as usual: `go get github.com/kavu/go-sigstats`.

Now you can use **go-phash**:

```go
package main

import "github.com/kavu/go-phash"

func main() {
	hash, _ := phash.ImageHash("~/my_cat.jpg")
	println(hash)
}
```
