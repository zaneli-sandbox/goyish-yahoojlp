package yahoojlp

func containStr(x string, ys ...string) bool {
	for _, y := range ys {
		if x == y {
			return true
		}
	}
	return false
}

func containInt(x int, ys ...int) bool {
	for _, y := range ys {
		if x == y {
			return true
		}
	}
	return false
}
