package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	testCases := []struct {
		x, y, want int
	}{
		{math.MinInt, 1, math.MinInt},
		{1, math.MinInt, 1},
		{-10, 6, 1000000},
		{-10, -6, 1000000},
		{-10, 5, -100000},
		{-10, -5, -100000},
		{-1, 1, -1},
		{-1, 0, 1},
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 2, 4},
		{2, 3, 8},
		{3, 3, 27},
		{math.MaxInt, 1, math.MaxInt},
		{1, math.MaxInt, 1},
		{-2, 63, math.MinInt},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: x:%d, y:%d", tC.x, tC.y)
		t.Run(testName, func(t *testing.T) {
			got, gotError := Pow(tC.x, tC.y)

			checkResults(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		x, y      int
		wantError error
	}{
		{"The result of math.MaxInt^2 is too large to store in an int", math.MaxInt, 2, ErrOverflowDetected},
		{"The result of 2^70 is too large to store in an int", 2, 70, ErrOverflowDetected},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := Pow(tC.x, tC.y)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func ExamplePow() {
	m, n := 3, 5
	result, err := Pow(m, n)
	if err != nil {
		fmt.Printf("Error calculating %d^%d: %v", m, n, err)
	} else {
		fmt.Printf("%d^%d = %d", m, n, result)
	}

	// Output: 3^5 = 243
}

func BenchmarkPow(b *testing.B) {
	for b.Loop() {
		_, err := Pow(3, 5)
		if err != nil {
			b.Fatal(err)
		}
	}
}
