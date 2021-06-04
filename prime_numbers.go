// Copied and adapted from: tinyurl.com/gosieve
// https://youtu.be/f6kdp27TYZs

// A concurrent prime sieve

package maths

import "context"

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

// GetPrimeNumbers returns a channel from which to siphon off the prime numbers in order, as needed.
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

// GetPrimeNumbersBelow fills a channel with the prime numbers below |n|.
func GetPrimeNumbersBelow(n int) <-chan int {
	n = Abs(n)

	primeChannel := make(chan int)
	go func() {
		if n < 3 {
			close(primeChannel)
			return
		}
		// Make slice ready for primes, starting from 2.
		numbers := make([]int, n-2)
		for ind := range numbers {
			numbers[ind] = ind + 2
		}

		// Euclidean sieve.
		for ind, val := range numbers {
			if val != 1 {
				primeChannel <- val
				for j := ind + val; j < n-2; j += val {
					if numbers[j] != 1 {
						numbers[j] = 1
					}
				}
			}
		}

		close(primeChannel)
	}()

	return primeChannel
}
