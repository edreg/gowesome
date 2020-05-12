package module01

func GCD(a, b int) int {

	for b != 0 {
		h := a % b
		a = b
		b = h
		//oder
		//a, b = b, a % b
	}

	return a
}
