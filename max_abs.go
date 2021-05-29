package maths

// Max returns the maximum int from a group of integers.
// Max() = 0
func Max(numbers ...int) int {
	switch len(numbers) {
	case 0:
		return 0
	case 1:
		return numbers[0]
	default:
		return max(numbers...)
	}
}

func max(numbers ...int) int {
	if len(numbers) > 2 {
		return max(numbers[0], max(numbers[1:]...))
	}

	if numbers[0] < numbers[1] {
		return numbers[1]
	}
	return numbers[0]
}

// Abs returns the absolute value of x.
// Does not handle math.MinInt64.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
