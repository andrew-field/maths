package maths

import (
	"fmt"
	"math/big"
)

// NumberOfDigits returns the number of digits of an integer or big.Int. Uses integer-string conversion.
func NumberOfDigits[T int | *big.Int](x T) int {
	s := fmt.Sprintf("%d", x) // Convert the number to a string.

	length := len(s)

	// If the number is negative, skip the '-' sign.
	if s[0] == '-' {
		return length - 1
	}

	return length
}

// GetDigits returns a slice filled with the digits of x in the same order (starting with the largest magnitude numbers, left to right).
func GetDigits[T int | *big.Int](x T) []int {
	digits := make([]int, 0)
	s := fmt.Sprintf("%d", x) // Convert the number to a string.

	// If the number is negative, skip the '-' sign.
	if s[0] == '-' {
		s = s[1:]
	}

	for _, val := range s {
		digits = append(digits, int(val-'0')) // Convert each character to its integer value by subtracting the ASCII value of '0'.
	}

	return digits
}
