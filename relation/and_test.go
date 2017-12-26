package relation

import "testing"

func TestAnd_Evaluate(t *testing.T) {
	a := And{}
	s := a.Evaluate()
	if s < 0 || s > 3 {
		t.Error("status is wrong")
	}
}
