package node

type ServiceNode struct {
	HostNode
	serviceName string
}

func NewServiceNode(hostName string, serviceName string, status uint8) *ServiceNode {
	n := ServiceNode{HostNode{ CommonNode{Status:status},hostName}, serviceName}
	return &n
}