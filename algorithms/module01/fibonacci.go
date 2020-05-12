package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	fibZero := 0
	fibOne := 1
	currentResult := 1
	predecessor := 1
	prePredecessor := 0

	if n == 0 {
		return fibZero
	}

	if n == 1 {
		return fibOne
	}

	for i := 2; i <= n; i++ {
		currentResult = predecessor + prePredecessor
		prePredecessor = predecessor
		predecessor = currentResult
	}

	return currentResult
}
