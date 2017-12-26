package node

import "testing"

func TestNewHostNode(t *testing.T) {
	hName := "test"
	hStatus := uint8(1)
	n := NewHostNode(hName, hStatus)
	if n.hostname != hName {
		t.Error("hostname was not set")
	}
	if n.Status != hStatus {
		t.Error("status was not set")
	}
}