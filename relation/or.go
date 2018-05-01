package relation

func EvaluateOr(childrenStatus []uint8) uint8 {
	status := uint8(3)
	for _, childStatus := range childrenStatus {
		if childStatus == 0 {
			return 0
		} else {
			if childStatus <= status {
				status = childStatus
			}
		}
	}
	return status
}