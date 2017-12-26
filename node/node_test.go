package node

import "testing"

func TestCommonNode_SetStatus(t *testing.T) {
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

func TestCommonNode_GetStatus(t *testing.T) {
	n := CommonNode{0}
	s := n.GetStatus()
	if s < 0 || s == nil {
		t.Errorf("status is %v", s)
	}
}