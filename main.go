package gomath

func division(a int, b int) (int, int) {
	var m int = mod(a, b)
	return (a - b) / m, m
}

func mod(a int, b int) int {
	if a > b {
		return mod(a-b, b)
	}
	return a
}
