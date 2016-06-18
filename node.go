package nue

import "errors"

const (
	nodeKindRoot = iota
	nodeKindInternal
	nodeKindLeaf
)

type Node struct {
	kind       int
	key        string
	value      *Route
	childNodes []*Node
}

func (n *Node) addChildNode(key string, value *Route, kind int) *Node {
	childNode := &Node{
		kind:  kind,
		key:   key,
		value: value,
	}
	n.childNodes = append(n.childNodes, childNode)
	return childNode
}

func (n *Node) getChildNode(key string) *Node {
	for _, node := range n.childNodes {
		if node.key == key {
			return node
		}
	}
	return nil
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
