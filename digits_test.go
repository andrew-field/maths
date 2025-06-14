package maths

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"testing"
)

func TestNumberOfDigits(t *testing.T) {
	testCasesInt := []struct {
		input, expectedResult int
	}{
		{math.MinInt, 19},
		{-10, 2},
		{-9, 1},
		{-1, 1},
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{500, 3},
		{4563198, 7},
		{math.MaxInt, 19},
	}

	for _, tC := range testCasesInt {
		testName := fmt.Sprintf("Input int: %d", tC.input)
		checkNumberOfDigits(tC.input, testName, tC.expectedResult, t)
	}

	testCasesBigInt := []struct {
		input          *big.Int
		expectedResult int
	}{
		{new(big.Int).Exp(big.NewInt(-10), big.NewInt(35), nil), 36},
		{big.NewInt(math.MinInt), 19},
		{big.NewInt(-10), 2},
		{big.NewInt(-9), 1},
		{big.NewInt(-1), 1},
		{big.NewInt(0), 1},
		{big.NewInt(1), 1},
		{big.NewInt(9), 1},
		{big.NewInt(10), 2},
		{big.NewInt(99), 2},
		{big.NewInt(100), 3},
		{big.NewInt(500), 3},
		{big.NewInt(4563198), 7},
		{big.NewInt(math.MaxInt), 19},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), 21},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(35), nil), 36},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(98), nil), 99},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(99), nil), 100},
	}

	for _, tC := range testCasesBigInt {
		testName := fmt.Sprintf("Input big.Int: %d", tC.input)
		checkNumberOfDigits(tC.input, testName, tC.expectedResult, t)
	}
}

func checkNumberOfDigits[T int | *big.Int](input T, testName string, expectedResult int, t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		if actualNumberOfDigits := NumberOfDigits(input); actualNumberOfDigits != expectedResult {
			t.Errorf("Actual number of digits: %d. Expected number of digits: %d", actualNumberOfDigits, expectedResult)
		}
	})
}

func TestGetDigits(t *testing.T) {
	testCasesInt := []struct {
		input          int
		expectedDigits []int
	}{
		{math.MinInt, []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}},
		{-10, []int{1, 0}},
		{-9, []int{9}},
		{-1, []int{1}},
		{0, []int{0}},
		{1, []int{1}},
		{9, []int{9}},
		{10, []int{1, 0}},
		{99, []int{9, 9}},
		{100, []int{1, 0, 0}},
		{500, []int{5, 0, 0}},
		{4563198, []int{4, 5, 6, 3, 1, 9, 8}},
		{math.MaxInt, []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}},
	}

	for _, tC := range testCasesInt {
		testName := fmt.Sprintf("Input int: %d", tC.input)
		checkDigits(tC.input, testName, tC.expectedDigits, t)
	}

	testCasesBigInt := []struct {
		input          *big.Int
		expectedDigits []int
	}{
		{new(big.Int).Exp(big.NewInt(-10), big.NewInt(21), nil), []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{big.NewInt(math.MinInt), []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}},
		{big.NewInt(-10), []int{1, 0}},
		{big.NewInt(-9), []int{9}},
		{big.NewInt(-1), []int{1}},
		{big.NewInt(0), []int{0}},
		{big.NewInt(1), []int{1}},
		{big.NewInt(9), []int{9}},
		{big.NewInt(10), []int{1, 0}},
		{big.NewInt(99), []int{9, 9}},
		{big.NewInt(100), []int{1, 0, 0}},
		{big.NewInt(500), []int{5, 0, 0}},
		{big.NewInt(4563198), []int{4, 5, 6, 3, 1, 9, 8}},
		{big.NewInt(math.MaxInt), []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil), []int{1, 2, 6, 7, 6, 5, 0, 6, 0, 0, 2, 2, 8, 2, 2, 9, 4, 0, 1, 4, 9, 6, 7, 0, 3, 2, 0, 5, 3, 7, 6}},
	}

	for _, tC := range testCasesBigInt {
		testName := fmt.Sprintf("Input big.Int: %d", tC.input)
		checkDigits(tC.input, testName, tC.expectedDigits, t)
	}
}

func checkDigits[T int | *big.Int](input T, testName string, expectedResult []int, t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		if digits := GetDigits(input); !slices.Equal(digits, expectedResult) {
			t.Errorf("Actual digits: %v Expected digits: %v", digits, expectedResult)
		}
	})
}

func ExampleNumberOfDigits() {
	ints := []int{12345, -12345, math.MaxInt, math.MinInt, 0, 1000000000, -1000000000, 99999999999999}

	for _, v := range ints {
		fmt.Printf("Number of digits in %d: %d\n", v, NumberOfDigits(v))
	}
	fmt.Println()

	bigInts := []*big.Int{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil)}
	for _, v := range bigInts {
		fmt.Printf("Number of digits in %d: %d\n", v, NumberOfDigits(v))
	}

	// Output:
	// Number of digits in 12345: 5
	// Number of digits in -12345: 5
	// Number of digits in 9223372036854775807: 19
	// Number of digits in -9223372036854775808: 19
	// Number of digits in 0: 1
	// Number of digits in 1000000000: 10
	// Number of digits in -1000000000: 10
	// Number of digits in 99999999999999: 14
	//
	// Number of digits in 100000000000000000000: 21
	// Number of digits in 1267650600228229401496703205376: 31
}

func ExampleGetDigits() {
	ints := []int{12345, -12345, math.MaxInt, math.MinInt, 0}

	for _, v := range ints {
		digits := GetDigits(v)
		fmt.Printf("The last digit of %d is %d\n", v, digits[len(digits)-1])
	}
	fmt.Println()

	bigInts := []*big.Int{new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil), new(big.Int).Exp(big.NewInt(3), big.NewInt(100), nil)}
	for _, v := range bigInts {
		digits := GetDigits(v)
		fmt.Printf("The last digit of %d is %d\n", v, digits[len(digits)-1])
	}

	// Output:
	// The last digit of 12345 is 5
	// The last digit of -12345 is 5
	// The last digit of 9223372036854775807 is 7
	// The last digit of -9223372036854775808 is 8
	// The last digit of 0 is 0
	//
	// The last digit of 1267650600228229401496703205376 is 6
	// The last digit of 515377520732011331036461129765621272702107522001 is 1
}

func BenchmarkNumberOfDigits(b *testing.B) {
	for b.Loop() {
		NumberOfDigits(12345678912345)
	}
}

func BenchmarkGetDigits(b *testing.B) {
	for b.Loop() {
		GetDigits(12345678912345)
	}
}

func FuzzNumberOfDigits(f *testing.F) {
	testCases := []int{
		math.MinInt, -10, -9, -1, 0, 1, 9, 10, 99, 100, 500, 4563198, math.MaxInt,
	}

	for _, v := range testCases {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, input int) {
		if input == 0 {
			t.Skip()
		}
		n := NumberOfDigits(input)

		orig := input
		for range n - 1 {
			input /= 10
		}
		if input == 0 || input/10 != 0 {
			t.Errorf("Number of digits for %d is incorrect: got %d", orig, n)
		}
	})
}

func FuzzGetDigits(f *testing.F) {
	testCases := []int{
		math.MinInt, -10, -9, -1, 0, 1, 9, 10, 99, 100, 500, 4563198, math.MaxInt,
	}

	for _, v := range testCases {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, input int) {
		digits := GetDigits(input)

		sum := 0
		exponent := 1
		for i := len(digits) - 1; i >= 0; i-- {
			sum += exponent * digits[i]
			exponent *= 10
		}

		if input < 0 {
			sum = -sum
		}

		if sum != input {
			t.Errorf("Digits of %d do not reconstruct the original number: got: %v, sum: %d", input, digits, sum)
		}
	})
}
