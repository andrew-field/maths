package maths

// Max returns the maximum int from a group of integers.
// Max() = 0
func Max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	return max(numbers...)
}

func max(numbers ...int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	if numbers[0] > numbers[1] {
		numbers[1] = numbers[0]
	}
	return max(numbers[1:]...)
}

// Abs returns the |x|.
// Does not handle math.MinInt64.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
