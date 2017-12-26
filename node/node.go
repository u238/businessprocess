package node

type Node interface {
	SetStatus(status uint8) (uint8, error)
}

type CommonNode struct {
	status uint8
}

func (n *CommonNode) SetStatus(status uint8) (uint8, error) {
	if n.status != status {
		n.status = status
		return 1, nil
	}
	return 0, nil
}