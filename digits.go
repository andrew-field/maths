package maths

import (
	"math/big"
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

// GetDigits returns and fills a channel with the digits of x
// starting with the smallest magnitude numbers (right to left).
func GetDigits(x int) <-chan int {
	digitsCh := make(chan int)

	go func() {
		if x < 0 {
			digitsCh <- -(x % 10) // Can successfully handle math.MinInt
			x /= -10
		}

		// 456/10 = 45 with int.
		for x > 0 {
			digitsCh <- x % 10
			x /= 10
		}

		close(digitsCh)
	}()

	return digitsCh
}

// GetDigitsBig fills and returns a channel with the digits of x
// starting with the smallest magnitude numbers (right to left).
func GetDigitsBig(x *big.Int) <-chan int {
	digitsCh := make(chan int)

	go func() {
		// Uses a new variable, altNumber, so as to not change the original number.
		altNumber := new(big.Int).Set(x)

		// Make number positive.
		altNumber.Abs(altNumber)

		ten := big.NewInt(10)
		var digit big.Int

		// Dividing these Ints by 10 truncates the decimal places.
		for altNumber.Sign() == 1 { // For altNumber > 0
			altNumber.QuoRem(altNumber, ten, &digit) // Go has this handy function. Sets altNumber to altNumber / 10 and sets digit to altNumber mod 10 (the last digit).
			digitsCh <- int(digit.Int64())
		}

		close(digitsCh)
	}()

	return digitsCh
}
