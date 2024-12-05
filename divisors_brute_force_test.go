package maths

import (
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
