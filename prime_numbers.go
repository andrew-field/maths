// Copied and adapted from: tinyurl.com/gosieve
// https://youtu.be/f6kdp27TYZs

// A concurrent prime sieve

package maths

import (
	"context"
	"math"
)

// GetPrimeNumbersBelowAndIncluding fills a channel with the prime numbers below and including |n|, in order. Uses a euclidean sieve.
// Cancel the context when the calling function has finished with the return values and does not care about further possible values.
func GetPrimeNumbersBelowAndIncluding(n int, ctx context.Context) <-chan int {
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

		maxPrime := int(math.Sqrt(float64(n)))
		// Step 1: All composite numbers below and including n must have a prime factor p such that p <= SQRT(n).
		// Hence to find all composite numbers, and therefore all prime numbers, generate all primes up to SQRT(n).
		smallPrimes := getPrimesUpTo(maxPrime)

		// Split the range [2, n] into numSegments equal (or nearly equal) segments.
		// Each segment is processed separately, reducing the maximum memory usage at any point.
		// Memory usage is roughly proportional to n / numSegments.
		// An easy way to implement the number of segments is to have it proportional (in this case set) to maxPrime.
		numSegments := maxPrime
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

			// Step 3: Mark composites within the segment using smallPrimes
			for _, p := range smallPrimes {
				// Find the minimum number in [start, end) that is a multiple of p
				minMultiple := ((start + p - 1) / p) * p
				if minMultiple < p*p {
					minMultiple = p * p
				}
				// Mark all multiples of p within the segment.
				for j := minMultiple; j < end; j += p {
					isComposite[j-start] = true
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
	isComposite := make([]bool, n+1)
	primes := []int{}

	for i := 2; i <= n; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			if i*i <= n {
				for j := i * i; j <= n; j += i {
					isComposite[j] = true
				}
			}
		}
	}
	return primes
}

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
