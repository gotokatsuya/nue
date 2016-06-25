# Nue - 鵺

Nue is a lightweight high performance HTTP request router for Golang.


## Build status

[![wercker status](https://app.wercker.com/status/6135ebbc86ffbe8fc6b370f18241bbea/s/master "wercker status")](https://app.wercker.com/project/bykey/6135ebbc86ffbe8fc6b370f18241bbea)


## Installation

```bash
$ go get github.com/gotokatsuya/nue
```


## Usage

```go
package main

import (
	"net/http"
	
	"github.com/gotokatsuya/nue"
)

func main() {
	handler := nue.New()
	handler.AddHandler("/hello", "/world", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	handler.AddNotFoundHandler(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Not found route."))
	})
	http.ListenAndServe(":8080", handler)
}
```
