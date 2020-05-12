package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	result := 0.0
	factor := 0.0
	pow := 0.0
	length := len(value)

	for i := 1; i <= length; i++ {
		factor = convertBaseStringToNumber(string(value[length-i]))
		pow = math.Pow(float64(base), float64(i-1))
		result += factor * pow
	}

	return int(result)
}

func convertBaseStringToNumber(n string) float64 {
	var result int
	switch {
	case n == "F":
		result = 15
	case n == "E":
		result = 14
	case n == "D":
		result = 13
	case n == "C":
		result = 12
	case n == "B":
		result = 11
	case n == "A":
		result = 10
	default:
		result, _ = strconv.Atoi(n)
	}

	return float64(result)
}
