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
// It will return an error if the concatenated digits exceed the range of an int. It will ignore any negative signs in the integers provided.
// If no integers are provided, it returns 0, nil.
func DigitsToInt(x ...int) (int, error) {
	if len(x) == 0 {
		return 0, nil
	}
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
// It will ignore any negative signs in the integers provided. If no integers are provided, it returns 0, big.NewInt(0).
func DigitsToBigInt(x ...int) *big.Int {
	if len(x) == 0 {
		return big.NewInt(0)
	}
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
