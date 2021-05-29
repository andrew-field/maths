package maths

// Factorial returns the factorial of |n|.
// It does not handle int overflows when numbers get too large. Use big.MulRange() instead.
func Factorial(n int) int {
	return fact(Abs(n))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Binomial returns to binomial coefficient of (|n|, |k|).
// It does not handle int overflows when numbers get too large. Use big.Binomial() instead.
func Binomial(n, k int) int {
	return (Factorial(n) / Factorial(k)) / Factorial(Abs(n)-Abs(k))
}
