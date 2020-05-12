package module01

import (
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var stringBuilder strings.Builder
	var res string

	for _, ch := range word {
		stringBuilder.WriteRune(ch)
		res = string(ch) + res
	}

	return res
}
