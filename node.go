package nue

import "errors"

const (
	nodeKindRoot = iota
	nodeKindInternal
	nodeKindLeaf
)

type Node struct {
	kind       int
	value      *Route
	childNodes map[string]*Node
}

func (n *Node) addChildNode(key string, value *Route, kind int) *Node {
	childNode := &Node{
		kind:       kind,
		value:      value,
		childNodes: make(map[string]*Node),
	}
	n.childNodes[key] = childNode
	return childNode
}

func (n *Node) getChildNode(key string) *Node {
	return n.childNodes[key]
}

func (n *Node) insertRoute(key string, value *Route) error {
	rootNode := n
	childNode := rootNode.getChildNode(key)
	if childNode == nil {
		rootNode = rootNode.addChildNode(key, nil, nodeKindInternal)
	} else {
		rootNode = childNode
	}
	rootNode.addChildNode("", value, nodeKindLeaf)
	return nil
}

var (
	errNodeNotFoundRoute = errors.New("Node: not found route.")
)

func (n *Node) findRoute(key string) (*Route, error) {
	rootNode := n
	childNode := rootNode.getChildNode(key)
	if childNode == nil {
		return nil, errNodeNotFoundRoute
	}
	rootNode = childNode
	leafNode := rootNode.getChildNode("")
	if leafNode != nil {
		return leafNode.value, nil
	}
	if slashNode := rootNode.getChildNode("/"); slashNode != nil {
		leafNode = slashNode.getChildNode("")
		if leafNode != nil {
			return leafNode.value, nil
		}
	}
	return nil, errNodeNotFoundRoute
}
