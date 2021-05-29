package maths

// Fact(x) returns the factorial of |x|.
// It does not handle int overflows when numbers get too large. Use big.MulRange() instead.
func Fact(n int) int {
	n = Abs(n)
	return fact(n)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Binomial(n, k) returns to binomial coefficient of (|n|, |k|).
// It does not handle int overflows when numbers get too large. Use big.Binomial() instead.
func Binomial(n, k int) int {
	return (Fact(n) / Fact(k)) / Fact(Abs(n)-Abs(k))
}
