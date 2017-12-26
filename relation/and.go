package relation

import "github.com/u238/businessprocess/node"

func EvaluateAnd(children []node.Node) uint8 {
	status := uint8(0)
	for _, child := range children {
		if child.GetStatus() >= status {
			status = child.GetStatus()
		}
	}
	return status
}