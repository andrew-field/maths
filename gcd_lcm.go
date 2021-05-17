package maths

import "math/big"

// GCD returns the greatest common divisor of a group of integers.
// GCD() = 0
// GCD(a, 0) = |a|
// Does not handle math.MinInt64
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
// Does not handle int overflows if the numbers get too large. Use LCMBig.
func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	var lcmResult int
	if len(numbers) == 1 {
		lcmResult = numbers[0]
	} else {
		lcmResult = lcm(numbers...)
	}

	if lcmResult < 0 {
		return -lcmResult
	}

	return lcmResult
}

func lcm(numbers ...int) int {
	if len(numbers) > 2 {
		return lcm(numbers[0], lcm(numbers[1:]...))
	}

	return (numbers[0] / GCD(numbers...)) * numbers[1]
}

// LCM returns the least common multple of a group of integers.
// LCM() = 0
// LCM(a, 0) = |a|
func LCMBig(numbers ...*big.Int) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	var lcmResult *big.Int
	if len(numbers) == 1 {
		lcmResult = new(big.Int).Set(numbers[0])
	} else {
		lcmResult = lcmBig(numbers...)
	}

	return lcmResult.Abs(lcmResult)
}

func lcmBig(numbers ...*big.Int) *big.Int {
	if len(numbers) > 2 {
		return lcmBig(numbers[0], lcmBig(numbers[1:]...))
	}

	gcd := new(big.Int).GCD(nil, nil, numbers[0], numbers[1])
	x, y := new(big.Int).Set(numbers[0]), new(big.Int).Set(numbers[1])

	return y.Mul(y, x.Div(x, gcd))
}
