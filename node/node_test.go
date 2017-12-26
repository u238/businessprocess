package node

import (
	"testing"
	"github.com/u238/businessprocess/node"
)

func TestCommonNode_SetStatus(t *testing.T) {
	n := CommonNode{0,[]node.Node{}}

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
	n := CommonNode{0, []node.Node{}}
	s := n.GetStatus()
	if s < 0 || s > 3 	{
		t.Errorf("status is %v", s)
	}
}