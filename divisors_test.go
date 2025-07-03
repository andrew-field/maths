package maths

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"testing"
)

var numberOfDivisorsTestCases = []struct {
	input, expectedResult int
}{
	{math.MinInt, 64},
	{-3, 2},
	{-2, 2},
	{-1, 1},
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 2},
	{4, 3},
	{5, 2},
	{6, 4},
	{7, 2},
	{8, 4},
	{9, 3},
	{10, 4},
	{100, 9},
	{500, 12},
	{45664, 12},
	{7931265, 32},
	{math.MaxInt, 96},
}

var getDivisorsTestCases = []struct {
	input int
	want  []int
}{
	{-4, []int{1, 2, 4}},
	{-3, []int{1, 3}},
	{-2, []int{1, 2}},
	{-1, []int{1}},
	{0, []int{}},
	{1, []int{1}},
	{2, []int{1, 2}},
	{3, []int{1, 3}},
	{4, []int{1, 2, 4}},
	{5, []int{1, 5}},
	{6, []int{1, 2, 3, 6}},
	{7, []int{1, 7}},
	{8, []int{1, 2, 4, 8}},
	{9, []int{1, 3, 9}},
	{10, []int{1, 2, 5, 10}},
	{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
	{500, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}},
	{45664, []int{1, 2, 4, 8, 16, 32, 1427, 2854, 5708, 11416, 22832, 45664}},
	{7931265, []int{1, 3, 5, 15, 17, 19, 51, 57, 85, 95, 255, 285, 323, 969, 1615, 1637, 4845, 4911, 8185, 24555, 27829, 31103, 83487, 93309, 139145, 155515, 417435, 466545, 528751, 1586253, 2643755, 7931265}},
	{math.MaxInt, []int{1, 7, 49, 73, 127, 337, 511, 889, 2359, 3577, 6223, 9271, 16513, 24601, 42799, 64897, 92737, 172207, 299593, 454279, 649159, 649657, 1205449, 2097151, 3124327, 4544113, 4547599, 6769801,
		11777599, 21870289, 31252369, 31833193, 47388607, 47424961, 82443193, 82506439, 153092023, 218766583, 218934409, 331720249, 331974727, 577102351, 577545073, 859764727, 1531366081,
		1532540863, 2281422937, 2323823089, 3969050863, 4042815511, 6018353089, 6022970047, 10727786041, 15969960559, 15982211857, 27783356041, 27804669943, 42128471623, 42160790329, 60247241209,
		111789723913, 111875482999, 194483492287, 194632689601, 289740712999, 295125532303, 421730688463, 783128380993, 1362428827207, 2028184990993, 2029740905839, 2952114819241, 4398048608257,
		7651399633543, 14197294936951, 14208186340873, 20303320287433, 30786340257799, 53559797434801, 99457304386111, 142123242012031, 215504381804593, 374918582043607, 558552173248639, 994862694084217,
		1482142380982609, 2578521676503991, 3909865212740473, 10374996666878263, 18049651735527937, 27369056489183311, 72624976668147841, 126347562148695559, 188232082384791343, 1317624576693539401,
		9223372036854775807}},
	{math.MinInt + 1, []int{1, 7, 49, 73, 127, 337, 511, 889, 2359, 3577, 6223, 9271, 16513, 24601, 42799, 64897, 92737, 172207, 299593, 454279, 649159, 649657, 1205449, 2097151, 3124327, 4544113, 4547599, 6769801,
		11777599, 21870289, 31252369, 31833193, 47388607, 47424961, 82443193, 82506439, 153092023, 218766583, 218934409, 331720249, 331974727, 577102351, 577545073, 859764727, 1531366081,
		1532540863, 2281422937, 2323823089, 3969050863, 4042815511, 6018353089, 6022970047, 10727786041, 15969960559, 15982211857, 27783356041, 27804669943, 42128471623, 42160790329, 60247241209,
		111789723913, 111875482999, 194483492287, 194632689601, 289740712999, 295125532303, 421730688463, 783128380993, 1362428827207, 2028184990993, 2029740905839, 2952114819241, 4398048608257,
		7651399633543, 14197294936951, 14208186340873, 20303320287433, 30786340257799, 53559797434801, 99457304386111, 142123242012031, 215504381804593, 374918582043607, 558552173248639, 994862694084217,
		1482142380982609, 2578521676503991, 3909865212740473, 10374996666878263, 18049651735527937, 27369056489183311, 72624976668147841, 126347562148695559, 188232082384791343, 1317624576693539401,
		9223372036854775807}},
}

var getDivisorsErrorTestCases = []struct {
	desc      string
	input     int
	wantError error
}{
	{"|math.MinInt| is too large a divisor to hold in an int", math.MinInt, ErrAbsoluteValueOfMinInt},
}

var sumOfDivisorsTestCases = []struct {
	input, want int
}{
	{-3, 4},
	{-2, 3},
	{-1, 1},
	{0, 0},
	{1, 1},
	{2, 3},
	{3, 4},
	{4, 7},
	{5, 6},
	{6, 12},
	{7, 8},
	{8, 15},
	{9, 13},
	{10, 18},
	{100, 217},
	{500, 1092},
	{45664, 89964},
	{7931265, 14152320},
}

var sumOfDivisorsErrorTestCases = []struct {
	desc      string
	input     int
	wantError error
}{
	{"The sum of the divisors is too large to store in an int", math.MaxInt, ErrOverflowDetected},
}

func TestNumberOfDivisors(t *testing.T) {
	testNumberOfDivisorsHelper(t, NumberOfDivisors)
}

func testNumberOfDivisorsHelper(t *testing.T, f func(int) int) {
	for _, tC := range numberOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualResult := f(tC.input); actualResult != tC.expectedResult {
				t.Errorf("Expected number of divisors: %d, actual number of divisors: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestGetDivisors(t *testing.T) {
	testGetDivisorsHelper(t, GetDivisors)
}

func testGetDivisorsHelper(t *testing.T, f func(int) (<-chan int, error)) {
	for _, tC := range getDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			divisorCh, gotError := f(tC.input)

			// Check there was no error.
			if gotError != nil {
				t.Errorf("Got error but didn't want one. Error: %v", gotError)
			}

			var actualDivisors []int
			for div := range divisorCh {
				actualDivisors = append(actualDivisors, div)
			}
			slices.Sort(actualDivisors)

			// Check if the actual divisors match the expected divisors.
			if !slices.Equal(actualDivisors, tC.want) {
				t.Errorf("Actual divisors: %v. Expected divisors: %v.", actualDivisors, tC.want)
			}
		})
	}

	for _, tC := range getDivisorsErrorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := f(tC.input)

			if !errors.Is(gotError, tC.wantError) {
				t.Errorf("Got %v, want %v", gotError, tC.wantError)
			}
		})
	}
}

func TestSumOfDivisors(t *testing.T) {
	testSumOfDivisorsHelper(t, SumOfDivisors)
}

func testSumOfDivisorsHelper(t *testing.T, f func(int) (int, error)) {
	for _, tC := range sumOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := f(tC.input)

			// Check there was no error.
			if gotError != nil {
				t.Errorf("Got error but didn't want one. Error: %v", gotError)
			}

			// Check if the actual result matches the expected result.
			if got != tC.want {
				t.Errorf("Expected sum of divisors: %d, actual sum of divisors: %d", tC.want, got)
			}
		})
	}

	for _, tC := range sumOfDivisorsErrorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := f(tC.input)

			if !errors.Is(gotError, tC.wantError) {
				t.Errorf("Got %v, want %v", gotError, tC.wantError)
			}
		})
	}
}

func ExampleNumberOfDivisors() {
	n := 28
	numDivisors := NumberOfDivisors(n)
	fmt.Println("The number of divisors of", n, "is", numDivisors)

	// Output: The number of divisors of 28 is 6
}

func ExampleGetDivisors() {
	n := 28
	divCh, err := GetDivisors(n)
	if err != nil {
		fmt.Printf("Error calculating the divisors of %d: %v", n, err)
		return
	}

	divisors := make([]int, 0)
	for d := range divCh {
		divisors = append(divisors, d)
	}
	fmt.Println("The divisors of", n, "are", divisors)

	// Output: The divisors of 28 are [1 2 4 7 14 28]
}

func ExampleSumOfDivisors() {
	n := 28
	sumDivisors, err := SumOfDivisors(n)
	if err != nil {
		fmt.Printf("Error calculating the sum of the divisors of %d: %v", n, err)
	} else {
		fmt.Println("The sum of the divisors of", n, "is", sumDivisors)
	}

	n = 3598428716789018112
	sumDivisors, err = SumOfDivisors(n)
	if err != nil {
		fmt.Printf("Error calculating the sum of the divisors of %d: %v", n, err)
	} else {
		fmt.Println("The sum of the divisors of", n, "is", sumDivisors)
	}

	// Output:
	// The sum of the divisors of 28 is 56
	// Error calculating the sum of the divisors of 3598428716789018112: failed to calculate 444879189109555200 * 42. The result is too large to hold in an int variable: arithmetic overflow detected
}

func BenchmarkNumberOfDivisors(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				NumberOfDivisors(input)
			}
		})
	}
}

func BenchmarkGetDivisors(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				divCh, err := GetDivisors(input)
				if err != nil {
					b.Fatal(err)
				}
				for range divCh { // Just iterating through the channel to benchmark the function.
				}
			}
		})
	}
}

func BenchmarkSumOfDivisors(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				_, err := SumOfDivisors(input)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
