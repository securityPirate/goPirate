package connector

import (
	"fmt"
	"strconv"
	"strings"
)

//Node hold data ssh connection
type Node struct {
	Host string
	Port int
	User string
}

//NewNode function to create new nodes
func NewNode(s string) *Node {
	node := &Node{
		Host: s,
	}

	if parts := strings.Split(node.Host, "@"); len(parts) > 1 {
		node.User = parts[0]
		node.Host = parts[1]
	}

	if parts := strings.Split(node.Host, ":"); len(parts) > 1 {
		node.Host = parts[0]
		node.Port, _ = strconv.Atoi(parts[1])
	}

	return node
}

func (node *Node) String() string {
	return fmt.Sprintf("%s:%d", node.Host, node.Port)
}
