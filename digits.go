package maths

import (
	"math/big"
	"strconv"
)

// Some methods use string-integer conversion, some functions stick with mathematical operations.

// NumberOfDigits returns the number of digits an int has.
func NumberOfDigits(x int) int {
	if x < 0 {
		return len(strconv.Itoa(x)) - 1
	}

	return len(strconv.Itoa(x))
}

// NumberOfDigitsBig returns the number of digits a big.Int has.
func NumberOfDigitsBig(x *big.Int) int {
	if x.Sign() == -1 {
		return len(x.String()) - 1
	}
	return len(x.String())
}

// Digits returns and fills a channel with the digits of a number
// starting with the smallest magnitude numbers (right to left).
func Digits(number int) <-chan int {
	digitsCh := make(chan int)

	go func() {
		if number < 0 {
			digitsCh <- -(number % 10) // Handles number = math.MinInt64
			number /= -10
		}

		// 456/10 = 45 with int.
		for number > 0 {
			digitsCh <- number % 10
			number /= 10
		}

		close(digitsCh)
	}()

	return digitsCh
}

// DigitsBig fills and returns a channel with the digits of a big.Int number
// starting with the smallest magnitude numbers (right to left).
func DigitsBig(number *big.Int) <-chan int {
	digitsCh := make(chan int)

	go func() {
		// Uses altNumber so as to not change the original number.
		altNumber := new(big.Int).Set(number)
		altNumber.Abs(altNumber)

		ten := big.NewInt(10)
		var digit big.Int

		// Dividing these Ints by 10 truncates the decimal places.
		for altNumber.Sign() == 1 {
			altNumber.QuoRem(altNumber, ten, &digit) // Go has this handy function. Sets altNumber to altNumber / 10 and sets digit to altNumber mod 10 (the last digit).
			digitsCh <- int(digit.Int64())
		}

		close(digitsCh)
	}()

	return digitsCh
}
