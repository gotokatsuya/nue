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
