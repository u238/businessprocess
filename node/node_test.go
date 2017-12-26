package node

import "testing"

func TestSetStatus(t *testing.T) {
	n := CommonNode{0}

	res, _ := n.SetStatus(0)
	if res != uint8(0) {
		t.Errorf("returned %d instead of 0", res)
	}

	res, _ = n.SetStatus(1)
	if res != uint8(1) {
		t.Errorf("returned %d instead of 1", res)
	}
}
