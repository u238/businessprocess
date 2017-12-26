package node

type HostNode struct {
	CommonNode
	hostname string
}

func NewHostNode(name string, status uint8) *HostNode {
	n := HostNode{CommonNode{Status:status}, name}
	return &n
}