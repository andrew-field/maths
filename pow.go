package maths

// Pow returns the x^|y|.
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}
	return pow(x, Abs(y), x)
}

func pow(x, y, product int) int {
	if y == 1 {
		return product
	}

	return pow(x, y-1, product*x)
}
