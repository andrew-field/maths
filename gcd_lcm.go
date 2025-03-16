package maths

import (
	"fmt"
	"math"
	"math/big"
	"slices"
)

// GCD returns the greatest common divisor of a group of integers i.e. the largest positive integer that divides each of the integers. This method uses the Euclidean algorithm.
// GCD() = GCD(0) = 0, nil.
// GCD(a, 0) = |a|, nil.
// If an overflow error is detected when  get too large, the function returns 0, ErrAbsoluteValueOfMinInt.
// In this case, use *bigInt.GCD() from the math/big package.
func GCD(numbers ...int) (int, error) {
	if len(numbers) == 0 || !slices.ContainsFunc(numbers, func(n int) bool { // If there are no numbers or no non-zero numbers (all zeroes), return 0, nil.
		return n != 0
	}) {
		return 0, nil
	}

	y := 0
	for _, x := range numbers {
		for y != 0 {
			x, y = y, x%y
		}
		y = x
	}

	absY, err := Abs(y) // A special case can not be made for math.MinInt because the GCD of just math.MinInt is |math.MinInt|, which can not be stored in an int variable.
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", y, err)
	}

	return absY, nil
}

// LCM returns the least common multiple of a group of integers i.e. the smallest positive integer that is divisible by each integer. This method uses GCD().
// LCM() = LCM(0, 0, ...) = 0, nil.
// LCM(a, 0) = 0, nil.
// If an overflow error is detected when the numbers get too large, the function returns 0, ErrOverflowDetected.
// In this case, use LCMBig.
func LCM(numbers ...int) (int, error) {
	// If there are no numbers or if one of the numbers is zero, return 0, nil.
	// There is an argument to be made that LCM(a, 0), for a != 0, should return an error. However, a decision has been made that it will return 0, nil instead.
	if len(numbers) == 0 || slices.ContainsFunc(numbers, func(n int) bool { // If there are no numbers or if one of the numbers is zero, return 0, nil.
		return n == 0
	}) {
		return 0, nil
	}

	lcmResult, err := Abs(numbers[0])
	if err != nil { // A special case can not be made for math.MinInt because the LCM of just math.MinInt is |math.MinInt|, which can not be stored in an int variable.
		return 0, fmt.Errorf("failed to get Abs(%d): %w", numbers[0], err)
	}
	if len(numbers) != 1 {
		for _, v := range numbers[1:] { // Calculate the LCM by repeatedly calculating the LCM of each pair of numbers.
			gcd, err := GCD(lcmResult, v)
			if err != nil {
				return 0, fmt.Errorf("failed to get GCD(%d, %d): %w", lcmResult, v, err)
			}

			// Check for overflow before multiplication.
			quotient := (lcmResult / gcd)
			// To make the overflow check valid, make sure v is positive at this stage.
			// The end result must be positive anyway, so by making sure this variable and therefore the lcmResult variable is positive, a final calculation taking the absolute value is not necessary.
			absV, err := Abs(v)
			if err != nil {
				return 0, fmt.Errorf("failed to get Abs(%d): %w", v, err)
			}
			if quotient > math.MaxInt/absV {
				return 0, fmt.Errorf("failed to calculate %d * %d. The result is too large to hold in an int variable: %w", quotient, absV, ErrOverflowDetected)
			}
			lcmResult = quotient * absV
		}
	}

	return lcmResult, nil // lcmResult can not be negative.
}

// LCMBig returns the least common multiple of a group of integers. This method uses *bigInt.GCD() from math/big.
// LCMBig() = LCM(0, 0, ...) = 0.
// LCM(a, 0) = 0.
func LCMBig(numbers ...*big.Int) *big.Int {
	zero := big.NewInt(0)
	// If there are no numbers or if one of the numbers is zero, return 0, nil.
	// There is an argument to be made that LCMBig(a, 0), for a != 0, should return an error. However, a decision has been made that it will return 0, nil instead.
	if len(numbers) == 0 || slices.ContainsFunc(numbers, func(n *big.Int) bool { // If there are no numbers or if one of the numbers is zero, return 0, nil.
		return n.Cmp(zero) == 0 // Can use == and != for *big.Int.
	}) {
		return zero
	}

	lcmResult := numbers[0]
	if len(numbers) != 1 {
		for _, v := range numbers[1:] { // Calculate the LCM by repeatedly calculating the LCM of each pair of numbers.
			gcd := new(big.Int).GCD(nil, nil, lcmResult, v) // Calculate the GCD.
			lcmResult.Mul(lcmResult, v.Div(v, gcd))         // v.Div sets v to the quotient but also returns v.
		}
	}

	return lcmResult.Abs(lcmResult)
}
