package maths

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"strconv"
	"testing"
)

func TestNumberOfDigits(t *testing.T) {
	testCasesInt := []struct {
		input, want int
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
		checkNumberOfDigits(tC.input, testName, tC.want, t)
	}

	testCasesBigInt := []struct {
		input *big.Int
		want  int
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
		checkNumberOfDigits(tC.input, testName, tC.want, t)
	}
}

func checkNumberOfDigits[T int | *big.Int](input T, testName string, expectedResult int, t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		if got := NumberOfDigits(input); got != expectedResult {
			t.Errorf("Actual number of digits: %d. Expected number of digits: %d", got, expectedResult)
		}
	})
}

func TestGetDigits(t *testing.T) {
	testCasesInt := []struct {
		input int
		want  []int
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
		checkDigits(tC.input, testName, tC.want, t)
	}

	testCasesBigInt := []struct {
		input *big.Int
		want  []int
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
		{new(big.Int).SetBit(new(big.Int), 100, 1), []int{1, 2, 6, 7, 6, 5, 0, 6, 0, 0, 2, 2, 8, 2, 2, 9, 4, 0, 1, 4, 9, 6, 7, 0, 3, 2, 0, 5, 3, 7, 6}},
	}

	for _, tC := range testCasesBigInt {
		testName := fmt.Sprintf("Input big.Int: %d", tC.input)
		checkDigits(tC.input, testName, tC.want, t)
	}
}

func checkDigits[T int | *big.Int](input T, testName string, expectedResult []int, t *testing.T) {
	t.Run(testName, func(t *testing.T) {
		if got := GetDigits(input); !slices.Equal(got, expectedResult) {
			t.Errorf("Actual digits: %v, Expected digits: %v", got, expectedResult)
		}
	})
}

func TestDigitsToInt(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{[]int{-9, -2, -2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}, math.MaxInt},
		{[]int{-1, 0}, 10},
		{[]int{-9}, 9},
		{[]int{-1}, 1},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{9}, 9},
		{[]int{1, 0}, 10},
		{[]int{9, 9}, 99},
		{[]int{1, 0, 0}, 100},
		{[]int{5, 0, 0}, 500},
		{[]int{4, 5, 6, 3, 1, 9, 8}, 4563198},
		{[]int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}, math.MaxInt},
		{[]int{123, 4567}, 1234567},
		{[]int{123, 4567, 8901}, 12345678901},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := DigitsToInt(tC.input...)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		input     []int
		wantError error
	}{
		{"The concatenation of the provided digits is too large to store in an int", []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}, strconv.ErrRange},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := DigitsToInt(tC.input...)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func TestDigitsToBigInt(t *testing.T) {
	testCases := []struct {
		input []int
		want  *big.Int
	}{
		{[]int{-1, -0, -0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, new(big.Int).Exp(big.NewInt(10), big.NewInt(21), nil)},
		{[]int{-9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}, new(big.Int).SetBit(new(big.Int), 63, 1)},
		{[]int{-1, 0}, big.NewInt(10)},
		{[]int{-9}, big.NewInt(9)},
		{[]int{-1}, big.NewInt(1)},
		{[]int{0}, big.NewInt(0)},
		{[]int{1}, big.NewInt(1)},
		{[]int{9}, big.NewInt(9)},
		{[]int{1, 0}, big.NewInt(10)},
		{[]int{9, 9}, big.NewInt(99)},
		{[]int{1, 0, 0}, big.NewInt(100)},
		{[]int{5, 0, 0}, big.NewInt(500)},
		{[]int{4, 5, 6, 3, 1, 9, 8}, big.NewInt(4563198)},
		{[]int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}, big.NewInt(math.MaxInt)},
		{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil)},
		{[]int{1, 2, 6, 7, 6, 5, 0, 6, 0, 0, 2, 2, 8, 2, 2, 9, 4, 0, 1, 4, 9, 6, 7, 0, 3, 2, 0, 5, 3, 7, 6}, new(big.Int).SetBit(new(big.Int), 100, 1)},
		{[]int{123, 4567}, big.NewInt(1234567)},
		{[]int{123, 4567, 8901}, big.NewInt(12345678901)},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			if got := DigitsToBigInt(tC.input...); got.Cmp(tC.want) != 0 {
				t.Errorf("Got: %v, want: %v", got, tC.want)
			}
		})
	}
}

func ExampleNumberOfDigits() {
	ints := []int{12345, -12345, math.MaxInt, math.MinInt, 0, 1000000000, -1000000000, 99999999999999}

	for _, v := range ints {
		fmt.Printf("Number of digits in %d: %d\n", v, NumberOfDigits(v))
	}
	fmt.Println()

	bigInts := []*big.Int{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), new(big.Int).SetBit(new(big.Int), 100, 1)}
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

	bigInts := []*big.Int{new(big.Int).SetBit(new(big.Int), 100, 1), new(big.Int).Exp(big.NewInt(3), big.NewInt(100), nil)}
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

func ExampleDigitsToInt() {
	digits := []int{1, 2, 3, 4, 5}

	result, err := DigitsToInt(digits...)
	if err != nil {
		fmt.Printf("Error calculating the integer from digits %v: %v\n", digits, err)
	} else {
		fmt.Printf("The integer from digits %v is %d\n", digits, result)
	}

	// Output:
	// The integer from digits [1 2 3 4 5] is 12345
}

func ExampleDigitsToBigInt() {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	result := DigitsToBigInt(digits...)

	fmt.Printf("The big integer from digits %v is %s\n", digits, result.String())

	// Output:
	// The big integer from digits [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20] is 1234567891011121314151617181920
}

func BenchmarkDigitsToInt(b *testing.B) {
	for b.Loop() {
		DigitsToInt(1, 2, 3, 4, 5)
	}
}

func BenchmarkDigitsToBigInt(b *testing.B) {
	for b.Loop() {
		DigitsToBigInt(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	}
}

func BenchmarkNumberOfDigits(b *testing.B) {
	for b.Loop() {
		NumberOfDigits(12345678912345)
	}
}

var testCases = []int{
	math.MinInt, -10, -9, -1, 0, 1, 9, 10, 99, 100, 500, 4563198, math.MaxInt,
}

func FuzzNumberOfDigits(f *testing.F) {
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
