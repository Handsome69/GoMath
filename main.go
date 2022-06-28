package gomath

func Division(a int, b int) (int, int) {
	var m int = Mod(a, b) //lol
	return (a - b) / m, m
}

func Mod(a int, b int) int {
	if a > b {
		return Mod(a-b, b)
	}
	return a
}
