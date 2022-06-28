package gomath

func Division(a int, b int) (int, int) {
	var m int = Mod(a, b)
	return (a - m) / b, m
}

func Mod(a int, b int) int {
	if a >= b {
		return Mod(a-b, b)
	}
	return a
}

func ModF(a float64, b float64) float64 {
	if a > b {
		return ModF(a-b, b)
	}
	return a
}

func Pow(a int, b int) int {
	var p int = 1
	for i := 0; i < b; i++ {
		p = p * a
	}
	return p
}

func GetDiv(a int) []int {
	var d []int
	for i := 1; i <= a; i++ {
		if Mod(a, i) == 0 {
			d = append(d, i)
		}
	}
	return d
}
