package maths

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
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

// DigitsToInt returns an int made from a concatenation of the given integers, in order.
func DigitsToInt(x ...int) (int, error) {
	var b strings.Builder

	for _, v := range x {
		s := strconv.Itoa(v)
		if s[0] == '-' {
			s = s[1:]
		}
		b.WriteString(s)
	}

	return strconv.Atoi(b.String())
}

// DigitsToBigInt returns a big.Int made from a concatenation of the given integers, in order.
func DigitsToBigInt(x ...int) *big.Int {
	var b strings.Builder

	for _, v := range x {
		s := strconv.Itoa(v)
		if s[0] == '-' {
			s = s[1:]
		}
		b.WriteString(s)
	}

	result, _ := new(big.Int).SetString(b.String(), 10)
	return result
}
