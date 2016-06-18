package nue

import (
	"net/http"
)

type Route struct {
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

func (r *Route) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	http.HandlerFunc(r.handler).ServeHTTP(rw, req)
}
