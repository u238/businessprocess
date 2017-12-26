package relation

import "github.com/u238/businessprocess/node"

type And struct {
	children []node.Node
}

func (n *And) Evaluate() uint8 {
	status := uint8(0)
	for _, child := range n.children {
		if child.GetStatus() >= status {
			status = child.GetStatus()
		}
	}
	return status
}