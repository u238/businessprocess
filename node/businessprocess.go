package node

import (
	rel "github.com/u238/businessprocess/relation"
	"errors"
)

type BusinessProcessNode struct {
	CommonNode
	name string
	evaluateRelation func([]uint8) uint8
	children []Node
}

func NewBusinessProcessNode(name string, relation string) (*BusinessProcessNode, error) {
	n := BusinessProcessNode{name: name}
	switch relation {
		case "AND":
			n.evaluateRelation = rel.EvaluateAnd
			break
		default:
			return nil, errors.New("relation type '" + relation + "' not implemented.")
	}
	return &n, nil
}
 func (n *BusinessProcessNode) GetChildrenStatus() []uint8 {
 	var statuses []uint8
 	for _, childStatus := range n.children {
 		statuses = append(statuses, childStatus.GetStatus())
	}
	return statuses
 }

func (n *BusinessProcessNode) EvaluateStatus() (uint8, error) {
	status := n.evaluateRelation(n.GetChildrenStatus())
	return n.SetStatus(status)
}