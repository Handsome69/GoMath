package gomath

func Division(a int, b int) (int, int) {
	var m int = Mod(a, b)
	return (a - b) / m, m
}

func Mod(a int, b int) int {
	if a > b {
		return Mod(a-b, b)
	}
	return a
}
