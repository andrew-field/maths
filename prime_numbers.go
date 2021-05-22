// Copied and adapted from: tinyurl.com/gosieve
// https://youtu.be/f6kdp27TYZs

// A concurrent prime sieve

package maths

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// GetPrimeNumbers returns a channel from which to siphon off the prime numbers in order, as needed.
// The prime sieve: Daisy-chain Filter processes.
func GetPrimeNumbers() <-chan int {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.

	primeCh := make(chan int) // Create return channel
	go func() {
		for {
			prime := <-ch
			primeCh <- prime
			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	}()

	return primeCh
}
