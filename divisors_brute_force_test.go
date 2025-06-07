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
	n := 28
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
	fmt.Println("Divisors of", n, "are", divisors)

	// Output: Divisors of 28 are [1 2 4 7 14 28]
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
