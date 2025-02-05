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

		segmentSize := n / numSegments

		for segment := 0; segment < numSegments; segment++ {
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
			wheel := []int{4, 2, 4, 2, 4, 6, 2, 6} // Wheel factorization for 2, 3, 5.
			for _, p := range smallPrimes {
				// Find the minimum number in [start, end) that is a multiple of p.
				minMultiple := ((start + p - 1) / p) * p
				if minMultiple < p*p {
					minMultiple = p * p
				}

				wIndex := 0
				// Mark all multiples of p within the segment.
				// Use the wheel to skip some multiples that are guaranteed to be composite.
				for j := minMultiple; j < end; {
					isComposite[j-start] = true
					j += p * wheel[wIndex]
					wIndex = (wIndex + 1) % 8 // len(wheel)
				}
			}

			// Step 4: Iterate over isComposite and send any number that is not marked as composite (i.e., false) to primeChannel.
			for i := 0; i < len(isComposite); i++ {
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
// Send a boolean to the Done channel when finished.
// The prime sieve: Daisy-chain Filter processes.
func GetPrimeNumbers() (<-chan int, chan<- bool) {
	ch := make(chan int) // Create a new channel.
	ctx, cancel := context.WithCancel(context.Background())
	go generate(ctx, ch) // Launch Generate goroutine.

	primeCh := make(chan int) // Create return channel.
	doneCh := make(chan bool) // Create done channel.
	go func() {
		<-doneCh
		cancel()
	}()

	go func() {
		for {
			prime := <-ch
			select {
			case primeCh <- prime:
			case <-ctx.Done():
				return
			}
			ch1 := make(chan int)
			go filter(ctx, ch, ch1, prime)
			ch = ch1
		}
	}()

	return primeCh, doneCh
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ctx context.Context, ch chan<- int) {
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
func filter(ctx context.Context, in <-chan int, out chan<- int, prime int) {
	for {
		var i int
		select {
		case i = <-in: // Receive value from 'in'.
		case <-ctx.Done():
			return
		}

		if i%prime != 0 {
			select {
			case out <- i: // Send 'i' to 'out'.
			case <-ctx.Done():
				return
			}
		}
	}
}
