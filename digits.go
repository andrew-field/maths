package maths

import (
	"math/big"
	"strconv"
)

// Some methods use string-integer conversion, some functions stick with mathematical operations.

// NumberOfDigitsOfInt returns the number of digits an int has.
func NumberOfDigitsOfInt(x int) int {
	if x < 0 {
		x = -x
	}

	return len(strconv.Itoa(x))
}

// NumberOfDigitsOfBigInt returns the number of digits a big.Int has.
func NumberOfDigitsOfBigInt(x *big.Int) int {
	if x.Sign() == -1 {
		return len(x.String()) - 1
	}
	return len(x.String())
}

// DigitsOfInt returns and fills a channel with the digits of a number
// starting with the smallest magnitude numbers (right to left).
func DigitsOfInt(number int) <-chan int {
	digitsChannel := make(chan int)

	go func() {
		if number < 0 {
			number = -number
		}

		// 456/10 = 45 with int.
		for number > 9 {
			digitsChannel <- number % 10
			number /= 10
		}

		digitsChannel <- number

		close(digitsChannel)
	}()

	return digitsChannel
}

// DigitsOfBigInt fills and returns a channel with the digits of a big.Int number
// starting with the smallest magnitude numbers (right to left).
func DigitsOfBigInt(number *big.Int) <-chan int {
	digitsChannel := make(chan int)

	go func() {
		// Uses altNumber so as to not change the original number.
		altNumber := new(big.Int).Set(number)
		altNumber.Abs(altNumber)

		ten := big.NewInt(10)
		var digit big.Int

		// Dividing these Ints by 10 truncates the decimal places so it is fine.
		for !altNumber.IsInt64() || altNumber.Int64() > 9 {
			// Go has this handy function. Sets altNumber to altNumber / 10 and sets digit to altNumber mod 10 (the last digit).
			altNumber.QuoRem(altNumber, ten, &digit)
			digitsChannel <- int(digit.Int64())
		}

		digitsChannel <- int(altNumber.Int64())

		close(digitsChannel)
	}()

	return digitsChannel
}
