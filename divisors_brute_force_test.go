package maths

import (
	"fmt"
	"slices"
	"testing"
)

func TestNumberOfDivisorsBruteForce(t *testing.T) {
	testNumberOfDivisorsHelper(t, NumberOfDivisorsBruteForce)
}

func TestGetDivisorsBruteForce(t *testing.T) {
	testGetDivisorsHelper(t, GetDivisorsBruteForce)
}

func TestSumOfDivisorsBruteForce(t *testing.T) {
	testSumOfDivisorsHelper(t, SumOfDivisorsBruteForce)
}

func ExampleNumberOfDivisorsBruteForce() {
	n := 28
	numDivisors := NumberOfDivisorsBruteForce(n)
	fmt.Println("Number of divisors of", n, "is", numDivisors)

	// Output: Number of divisors of 28 is 6
}

func ExampleGetDivisorsBruteForce() {
	n := 30
	divCh, err := GetDivisorsBruteForce(n)
	if err != nil {
		fmt.Printf("Error calculating the divisors of %d: %v", n, err)
		return
	}

	divisors := make([]int, 0)
	for d := range divCh {
		divisors = append(divisors, d)
	}
	slices.Sort(divisors)
	fmt.Printf("Divisors of %d are %v", n, divisors)

	// Output: Divisors of 30 are [1 2 3 5 6 10 15 30]
}

func ExampleSumOfDivisorsBruteForce() {
	n := 28
	sumDivisors, err := SumOfDivisorsBruteForce(n)
	if err != nil {
		fmt.Printf("Error calculating sum of the divisors of %d: %v", n, err)
	} else {
		fmt.Println("Sum of the divisors of", n, "is", sumDivisors)
	}

	n = 3598428716789018112
	sumDivisors, err = SumOfDivisorsBruteForce(n)
	if err != nil {
		fmt.Printf("Error calculating the sum of the divisors of %d: %v", n, err)
	} else {
		fmt.Println("Sum of the divisors of", n, "is", sumDivisors)
	}

	// Output:
	// Sum of the divisors of 28 is 56
	// Error calculating the sum of the divisors of 3598428716789018112: failed to calculate 9060329447629492072 + 399825412976557568. The result is too large to hold in an int variable: arithmetic overflow detected
}

var divisorBenchmarkInputs = []int{10, 100, 1000, 10000, 100000, 1000000}

func BenchmarkNumberOfDivisorsBruteForce(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				NumberOfDivisorsBruteForce(input)
			}
		})
	}
}

func BenchmarkGetDivisorsBruteForce(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				divCh, _ := GetDivisorsBruteForce(input)
				for range divCh {
					// Just iterating through the channel to benchmark the function.
				}
			}
		})
	}
}

func BenchmarkSumOfDivisorsBruteForce(b *testing.B) {
	for _, input := range divisorBenchmarkInputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				SumOfDivisorsBruteForce(input)
			}
		})
	}
}
