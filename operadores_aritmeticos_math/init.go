package operadores_aritmeticos_math

import "math"

func Resto() int {
	a := 10
	b := 3
	resto := a % b
	return resto
}

// GetPi return math pi
func GetPi() float64 {
	pi := math.Pi
	return pi
}

func GetPower(
	base float64,
	exp float64,
) float64 {
	power := math.Pow(base, exp)
	return power
}
