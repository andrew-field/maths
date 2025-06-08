package maths

import (
	"context"
	"math"
)

// GetPrimeNumbersBelowAndIncluding fills a channel with the prime numbers below and including |n|, in order. Uses a euclidean sieve.
// Cancel the context when the calling function has finished with the return values and does not care about further possible values.
func GetPrimeNumbersBelowAndIncluding(ctx context.Context, n int) <-chan int {
	// Special case when n is equal to math.MinInt.
	// In this case, getting the absolute value would return an error, but the prime numbers below |math.MinInt| and math.MaxInt are the same.
	// (i.e. math.MaxInt, 2⁶³ - 1, is not a prime itself).
	if n == math.MinInt {
		n = math.MaxInt
	}

	if n < 0 {
		n = -n // Because math.MinInt case is checked above, this can not panic with an error.
	}

	primeChannel := make(chan int)
	go func() {
		defer close(primeChannel)

		if n < 2 {
			return
		}

		// Step 1: All composite numbers below and including n must have a prime factor p such that p <= SQRT(n).
		// Hence to find all composite numbers, and therefore all prime numbers, generate all primes up to SQRT(n).
		maxPrime := int(math.Sqrt(float64(n)))
		smallPrimes := getPrimesUpTo(maxPrime)

		// Split the range [2, n] into numSegments equal (or nearly equal) segments.
		// Each segment is processed separately, reducing the maximum memory usage at any point.
		// Memory usage is roughly proportional to n / numSegments.
		// An easy way to implement the number of segments is to have it proportional to maxPrime.
		numSegments := maxPrime / 100
		if numSegments < 10 {
			numSegments = 1
		}

		segmentSize := (n - 1) / numSegments // n - 1 because that is the range or [2 n] (we are excluding 1).

		for segment := range numSegments {
			start := 2 + segment*segmentSize
			var end int
			if segment == numSegments-1 {
				end = n + 1
			} else {
				end = start + segmentSize
			}

			// Step 2: Create a slice, isComposite, for the current segment to track whether numbers in this segment are composite.
			isComposite := make([]bool, end-start)

			// Step 3: Mark composites within the segment using smallPrimes.
			for _, p := range smallPrimes {
				// Find the minimum number in [start, end) that is a multiple of p.
				minMultiple := max(((start+p-1)/p)*p, p*p)

				// Mark all multiples of p within the segment.
				for j := minMultiple; j < end; j += p {
					isComposite[j-start] = true
				}
			}

			// Step 4: Iterate over isComposite and send any number that is not marked as composite (i.e., false) to primeChannel.
			for i := range isComposite {
				if !isComposite[i] {
					select {
					case primeChannel <- start + i:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return primeChannel
}

func getPrimesUpTo(n int) []int {
	if n < 2 {
		return []int{}
	}
	if n <= 5 {
		switch n {
		case 2:
			return []int{2}
		case 3:
			return []int{2, 3}
		case 4, 5:
			return []int{2, 3, 5}
		}
	}

	// Wheel factorization for 2, 3, 5.
	wheel := []int{4, 2, 4, 2, 4, 6, 2, 6}
	wIndex := 0
	candidate := 7

	// isComposite[i] represents whether candidate + i is composite.
	isComposite := make([]bool, n+1)

	primes := []int{2, 3, 5}
	for candidate <= n {
		if !isComposite[candidate] {
			primes = append(primes, candidate)
			// Mark multiples of candidate as composite
			for j := candidate * candidate; j <= n; j += candidate {
				isComposite[j] = true
			}
		}
		candidate += wheel[wIndex]
		wIndex = (wIndex + 1) % 8 // len(wheel)
	}

	return primes
}

// Copied and adapted from: tinyurl.com/gosieve
// https://youtu.be/f6kdp27TYZs

// GetPrimeNumbers returns a channel from which to siphon off the prime numbers in order, as needed.
// Cancel the context when the calling function has finished with the return values and does not care about further possible values.
// The prime sieve: Daisy-chain Filter processes.
func GetPrimeNumbers(ctx context.Context) <-chan int {
	ch := make(chan int) // Initial channel to send all the numbers on.

	go generate(ctx, ch) // Launch Generate goroutine to send all the numbers.

	primeCh := make(chan int) // Create the main prime channel to return.

	go func() {
		defer close(primeCh)

		// The flow of numbers is: Generate function -> filter for '2' -> filter for '3' -> filter for '5' -> ... -> prime channel.
		for {
			prime := <-ch // Receive a prime number from the last 'out' channel. (Will succeed immediately is the channel is closed so it will not block.)
			select {
			case primeCh <- prime: // Send the prime number to the main prime channel. The main blocking operation.
			case <-ctx.Done():
				return
			}
			ch1 := make(chan int)     // Create a new 'out' channel.
			go filter(ch, ch1, prime) // Append a new filter to the end of the chain, using the input of the last filter and the new out channel.
			ch = ch1                  // Output of this filter is the input of the next filter.
		}
	}()

	return primeCh
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ctx context.Context, ch chan<- int) {
	defer close(ch)
	for i := 2; ; i++ {
		select {
		case ch <- i: // Send 'i' to channel 'ch'.
		case <-ctx.Done():
			return
		}
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	defer close(out)
	for i := range in {
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
