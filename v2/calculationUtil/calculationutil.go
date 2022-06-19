package calculationutil

func AddTwoNum(a, b int) int {
	return a + b
}

func SubBFromA(a, b int) int {
	return a - b
}

func AddNums(args ...int) int {
	s := 0
	for _, k := range args {
		s += k
	}
	return s
}
