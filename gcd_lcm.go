package maths

// GCD returns the greatest common divisor of a group of integers.
// GCD() = 0
// GCD(a, 0) = |a|
func GCD(numbers ...int) int {
	y := 0
	for _, x := range numbers {
		for y != 0 {
			x, y = y, x%y
		}
		y = x
	}

	if y < 0 {
		return -y
	}

	return y
}

// LCM returns the least common multple of a group of integers.
// LCM() = 0
// LCM(a, 0) = |a|
func LCM(numbers ...int) int {
	gcd := GCD(numbers...)

	if gcd == 0 {
		return 0
	}

	lcm := numbers[0] / gcd

	for i := 1; i < len(numbers); i++ {
		lcm *= numbers[i]
	}

	if lcm < 0 {
		return -lcm
	}

	return lcm
}
