package node

type Node interface {
	SetStatus(status uint8) (uint8, error)
}

type CommonNode struct {
	Status uint8
}

func (n *CommonNode) SetStatus(status uint8) (uint8, error) {
	if n.Status != status {
		n.Status = status
		return 1, nil
	}
	return 0, nil
}
