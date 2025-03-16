package maths

import (
	"math/big"
	"slices"
	"strconv"
)

// NumberOfDigits returns the number of digits of an integer. Uses integer-string conversion.
func NumberOfDigits(x int) int {
	if x < 0 { // Can not write x = -x because x is an int and math.MinInt is not representable as a positive int.
		return len(strconv.Itoa(x)) - 1
	}

	return len(strconv.Itoa(x))
}

// NumberOfDigitsBig returns the number of digits of a big.Int. Uses integer-string conversion.
func NumberOfDigitsBig(x *big.Int) int {
	if x.Sign() == -1 {
		return len(x.String()) - 1
	}
	return len(x.String())
}

// GetDigits returns a slice filled with the digits of x in the same order (starting with the largest magnitude numbers, left to right).
func GetDigits(x int) []int {
	digits := make([]int, 0)
	if x == 0 {
		digits = append(digits, 0)
	}

	if x < 0 {
		digits = slices.Insert(digits, 0, -(x % 10)) // Can successfully handle math.MinInt
		x /= -10
	}

	// 456/10 = 45 with int.
	for x > 0 {
		digits = slices.Insert(digits, 0, x%10)
		x /= 10
	}

	return digits
}

// GetDigitsBig returns a slice filled with the digits of x in the same order (starting with the largest magnitude numbers, left to right).
func GetDigitsBig(x *big.Int) []int {
	digits := make([]int, 0)
	if x.Sign() == 0 {
		digits = append(digits, 0)
	}

	// Uses a new variable, altNumber, so as to not change the original number.
	altNumber := new(big.Int)

	// Make number positive.
	altNumber.Abs(x)

	ten := big.NewInt(10)
	var digit big.Int

	for altNumber.Sign() == 1 { // For altNumber > 0.
		altNumber.QuoRem(altNumber, ten, &digit) // Go has this handy function. Sets altNumber to altNumber / 10 and sets digit to altNumber mod 10 (the last digit).
		digits = slices.Insert(digits, 0, int(digit.Int64()))
	}

	return digits
}
