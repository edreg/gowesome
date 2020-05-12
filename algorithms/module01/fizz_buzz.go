package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		isDivByThree := i%3 == 0
		isDivByFive := i%5 == 0
		index := i - 1

		if isDivByThree && isDivByFive {
			result[index] = "Fizz Buzz"
		} else if isDivByThree {
			result[index] = "Fizz"
		} else if isDivByFive {
			result[index] = "Buzz"
		} else {
			result[index] = strconv.Itoa(i)
		}
	}

	fmt.Print(strings.Join(result, ", "))
	fmt.Println()
}
