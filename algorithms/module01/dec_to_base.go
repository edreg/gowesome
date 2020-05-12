package module01

import (
	"math"
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//

func DecToBase(dec, base int) string {
	result := ""

	mod, rounded := getModAndRounded(dec, base)
	result = convertNumberToBaseString(mod) + result

	for rounded > 0 {
		mod, rounded = getModAndRounded(rounded, base)
		result = convertNumberToBaseString(mod) + result
	}

	return result
}

func getModAndRounded(dec int, base int) (int, int) {
	mod := dec % base
	rounded := math.Floor(float64(dec / base))
	return mod, int(rounded)
}

func convertNumberToBaseString(n int) string {
	var result string
	switch {
	case n == 15:
		result = "F"
	case n == 14:
		result = "E"
	case n == 13:
		result = "D"
	case n == 12:
		result = "C"
	case n == 11:
		result = "B"
	case n == 10:
		result = "A"
	default:
		result = strconv.Itoa(n)
	}

	return result
}
