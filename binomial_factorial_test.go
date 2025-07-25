package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input, want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{10, 3628800},
		{20, 2432902008176640000},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := Factorial(tC.input)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		input     int
		wantError error
	}{
		{"n must be non-negative", -1, ErrNegativeNumber},
		{"The result of 21! is too large to store in an int", 21, ErrOverflowDetected},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := Factorial(tC.input)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func TestBinomial(t *testing.T) {
	testCases := []struct {
		n, k, want int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 1, 2},
		{2, 2, 1},
		{10, 10, 1},
		{10, 5, 252},
		{math.MaxInt, math.MaxInt, 1},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: n:%d, k:%d", tC.n, tC.k)
		t.Run(testName, func(t *testing.T) {
			got, gotError := Binomial(tC.n, tC.k)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		n, k      int
		wantError error
	}{
		{"n must be non-negative", -10, 5, ErrNegativeNumber},
		{"k must be non-negative", 10, -5, ErrNegativeNumber},
		{"k must not be larger than n", 5, 6, ErrNLessThanK},
		{"The result of (70, 35) is too large to store in an int", 70, 35, ErrOverflowDetected},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := Binomial(tC.n, tC.k)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func ExampleFactorial() {
	n := 10
	p, err := Factorial(n)
	if err != nil {
		fmt.Printf("Error calculating the factorial of %d: %v\n", n, err)
	} else {
		fmt.Printf("The factorial of %d is %d\n", n, p)
	}

	n = 21
	p, err = Factorial(n)
	if err != nil {
		fmt.Printf("Error calculating the factorial of %d: %v\n", n, err)
	} else {
		fmt.Printf("The factorial of %d is %d\n", n, p)
	}

	// Output:
	// The factorial of 10 is 3628800
	// Error calculating the factorial of 21: the result of 21! is too large to hold in an int variable: arithmetic overflow detected
}

func ExampleBinomial() {
	n, k := 10, 3
	p, err := Binomial(n, k)
	if err != nil {
		fmt.Printf("Error calculating the binomial coefficient of %d choose %d: %v\n", n, k, err)
	} else {
		fmt.Printf("The binomial coefficient of %d choose %d is %d\n", n, k, p)
	}

	n, k = 70, 35
	p, err = Binomial(n, k)
	if err != nil {
		fmt.Printf("Error calculating the binomial coefficient of %d choose %d: %v\n", n, k, err)
	} else {
		fmt.Printf("The binomial coefficient of %d choose %d is %d\n", n, k, p)
	}

	// Output:
	// The binomial coefficient of 10 choose 3 is 120
	// Error calculating the binomial coefficient of 70 choose 35: the result of (70, 35) is too large to hold in an int variable: arithmetic overflow detected
}

func BenchmarkBinomial(b *testing.B) {
	for b.Loop() {
		_, err := Binomial(10, 3)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for b.Loop() {
		_, err := Factorial(10)
		if err != nil {
			b.Fatal(err)
		}
	}
}
