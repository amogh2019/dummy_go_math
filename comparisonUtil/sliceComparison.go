package comparisonUtil

func CompareSlices(a []int, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	// compare their lengths and values in order
	if len(a) != len(b) {
		return false
	}

	for idx, val := range a {
		if b[idx] != val {
			return false
		}
	}
	return true
}
