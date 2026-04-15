package maths

import "errors"

var (
	// ErrNegativeNumber is an error that indicates a negative number has been used where not permitted.
	ErrNegativeNumber = errors.New("number must be non-negative")

	// ErrOverflowDetected is an error that indicates an arithmetic overflow has been detected.
	ErrOverflowDetected = errors.New("arithmetic overflow detected")

	// ErrNLessThanK is an error that indicates an invalid binomial coefficient calculation.
	ErrNLessThanK = errors.New("to calculate n choose k, n must be larger than or equal to k")
)
