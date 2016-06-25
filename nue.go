package nue

import (
	"fmt"
	"net/http"
	"path"
)

type Nue struct {
	*Router
	notFoundRouteHandler func(http.ResponseWriter, *http.Request)
}

func New() *Nue {
	return &Nue{
		Router: NewRouter(),
	}
}

func (n *Nue) Add(prefix, pattern string, h func(http.ResponseWriter, *http.Request)) error {
	return n.addRoute(prefix, pattern, h)
}

func (n *Nue) AddNotFoundHandler(h func(http.ResponseWriter, *http.Request)) {
	n.notFoundRouteHandler = h
}

func (n *Nue) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix, pattern := splitURLPath(path.Clean(r.URL.Path))
	h, err := n.findRoute(prefix, pattern)
	if err != nil {
		n.notFoundRouteHandler(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

func (n *Nue) ShowNodes() {
	for key, node := range n.nodes {
		fmt.Printf("Group: %s\n", key)
		for key, _ := range node.childNodes {
			fmt.Printf("Node: %s\n", key)
		}
	}
}
