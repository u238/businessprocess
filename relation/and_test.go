package relation

import (
	"testing"
	"github.com/u238/businessprocess/node"
)

func TestEvaluateAnd(t *testing.T) {
	a := []node.Node{}
	s := EvaluateAnd(a)
	if s < 0 || s > 3 {
		t.Error("status is wrong")
	}
}
