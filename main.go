package gomath

type expression func(float64) float64

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

func IsPrime(a int) bool {
	return len(GetDiv(a)) < 3
}

func Sum(e expression, start int, end int) float64 {
	var result float64 = 0
	for i := start; i <= end; i++ {
		result += e(float64(i))
	}
	return result
}

func Prod(e expression, start int, end int) float64 {
	var result float64 = 1
	for i := start; i <= end; i++ {
		result *= e(float64(i))
	}
	return result
}

func Pi() float64 {
	return 3.141592653589793238462643
}

func E() float64 {
	return 2.718281828459045235360287
}
