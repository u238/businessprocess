package node

type Node interface {
	SetStatus(status uint8) (uint8, error)
	GetStatus() (uint8)
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

func (n *CommonNode) GetStatus() uint8 {
	return n.Status
}