# Warning! Original pHash library is unmaintained. Many users face build problems with it and this wrapper. You have been warned. 


# go-phash #

[![Build Status](https://travis-ci.org/kavu/go-phash.png?branch=master)](https://travis-ci.org/kavu/go-phash)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkavu%2Fgo-phash.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkavu%2Fgo-phash?ref=badge_shield)

**go-phash** is a simple [pHash](http://phash.org) wrapper library for the Go programming language.

You can view documentation on [godoc.org](http://godoc.org/github.com/kavu/go-phash "go-phash documentation").

## Example ##

First of all, install **go-phash** as usual: `go get github.com/kavu/go-phash`.

Now you can use **go-phash**:

```go
package main

import "github.com/kavu/go-phash"

func main() {
	hash, _ := phash.ImageHashDCT("~/my_cat.jpg")
	println(hash)
}
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkavu%2Fgo-phash.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkavu%2Fgo-phash?ref=badge_large)