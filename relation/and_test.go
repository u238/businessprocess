package relation

import (
	"testing"
)

func TestEvaluateAnd(t *testing.T) {
	a := []uint8{0, 1, 2}
	s := EvaluateAnd(a)
	if s < 0 || s > 3 {
		t.Error("status is wrong")
	}
}
