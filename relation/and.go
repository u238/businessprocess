package relation

func EvaluateAnd(childrenStatus []uint8) uint8 {
	status := uint8(0)
	for _, childStatus := range childrenStatus {
		if childStatus >= status {
			status = childStatus
		}
	}
	return status
}