package node

import "testing"

func TestNewServiceNode(t *testing.T) {
	hName := "testHost"
	sName := "testService"
	sStatus := uint8(1)
	n := NewServiceNode(hName, sName, sStatus)
	if n.hostname != hName {
		t.Error("hostname was not set")
	}
	if n.serviceName != sName {
		t.Error("serviceName was not set")
	}
	if n.Status != sStatus {
		t.Error("service status was not set")
	}
}
