package main

import (
	"github.com/gotokatsuya/nue"
	"net/http"
)

func main() {
	handler := nue.New()
	handler.Add("/hello", "/world", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", handler)
}
