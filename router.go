package nue

import (
	"errors"
	"net/http"
)

type Router struct {
	nodes map[string]*Node
}

func NewRouter() *Router {
	return &Router{
		nodes: make(map[string]*Node),
	}
}

func (r *Router) addRoute(prefix, path string, h func(http.ResponseWriter, *http.Request)) error {
	node := r.nodes[prefix]
	if node == nil {
		r.nodes[prefix] = &Node{kind: nodeKindRoot, childNodes: make(map[string]*Node)}
	}
	return r.nodes[prefix].insertRoute(path, &Route{path: path, handler: h})
}

var (
	errRouterNotFoundRoute = errors.New("Router: not found route.")
)

func (r *Router) findRoute(prefix, path string) (*Route, error) {
	node := r.nodes[prefix]
	if node != nil {
		return node.findRoute(path)
	}
	return nil, errRouterNotFoundRoute
}
