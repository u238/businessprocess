package node

import "testing"

func TestNewBusinessProcessNode(t *testing.T) {
	bpName := "test"
	bpn, err := NewBusinessProcessNode(bpName, "AND")
	if err != nil {
		t.Error("error creating businessprocess node: " + err.Error())
	}
	if bpn.name != bpName {
		t.Error("name was not set in businessprocess node")
	}
}
