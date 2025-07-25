package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	testCasesInt := []struct {
		input []int
		want  int
	}{
		{[]int{0}, 0},
		{[]int{-10}, -10},
		{[]int{10}, 10},
		{[]int{-10, -10}, -10},
		{[]int{-10, -9}, -9},
		{[]int{-1, 0}, 0},
		{[]int{0, -1}, 0},
		{[]int{0, 0}, 0},
		{[]int{0, 1}, 1},
		{[]int{-1, 1}, 1},
		{[]int{1, 10}, 10},
		{[]int{-100, -100, -100}, -100},
		{[]int{-100, -100, -99}, -99},
		{[]int{-100, -1, 0}, 0},
		{[]int{1, 100, 10}, 100},
		{[]int{10, 5, 7, 2, 3, 6, 8, 1, 4, 9}, 10},
	}

	for _, tC := range testCasesInt {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			if got := Max(tC.input...); got != tC.want {
				t.Errorf("Expected result: %d, got result: %d", tC.want, got)
			}
		})
	}

	testCasesFloat := []struct {
		input []float64
		want  float64
	}{
		{[]float64{0.0}, 0.0},
		{[]float64{-0.9999, -1}, -0.9999},
		{[]float64{0.9999, 1}, 1},
	}

	for _, tC := range testCasesFloat {
		testName := fmt.Sprintf("Input: %0.4f", tC.input)
		t.Run(testName, func(t *testing.T) {
			if got := Max(tC.input...); got != tC.want {
				t.Errorf("Expected result: %f, got result: %f", tC.want, got)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	testCases := []struct {
		input, want int
	}{
		{math.MinInt + 1, math.MaxInt},
		{-100, 100},
		{-1, 1},
		{0, 0},
		{1, 1},
		{100, 100},
		{math.MaxInt, math.MaxInt},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := Abs(tC.input)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		input     int
		wantError error
	}{
		{"The value of |math.MinInt| can not be stored as an int", math.MinInt, ErrAbsoluteValueOfMinInt},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := Abs(tC.input)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func ExampleAbs() {
	n := -10
	absValue, err := Abs(n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The absolute value of", n, "is", absValue)
	}

	// Output: The absolute value of -10 is 10
}

func ExampleMax() {
	maxValue := Max(-100, -5, 0, 5, 10)
	fmt.Println("The maximum value is", maxValue)

	// Output: The maximum value is 10
}
