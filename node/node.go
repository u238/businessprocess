package node

import "github.com/u238/businessprocess/node"

type Node interface {
	SetStatus(status uint8) (uint8, error)
	EvaluateStatus() (uint8, error)
	GetStatus() (uint8)
}

type CommonNode struct {
	Status uint8
	parents []node.Node
}

func (n *CommonNode) SetStatus(status uint8) (uint8, error) {
	if n.Status != status {
		n.Status = status
		for _, parent := range n.parents {
			return parent.EvaluateStatus()
		}
		return 1, nil
	}
	return 0, nil
}

func (n *CommonNode) EvaluateStatus() (uint8, error) {
	return n.Status, nil
}


func (n *CommonNode) GetStatus() uint8 {
	return n.Status
}