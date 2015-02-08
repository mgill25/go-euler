// Some stuff on primes
// +build primes

package main

import "fmt"

func sieveOfEratosthenes(n int) []int {
	// Space required: O(n)
	// Time required: Loose upper bound: O(n^2). But it is less than that.
	// We can be more precise here...
	// O(n * log(log(n)))
	sieve := make([]bool, n)
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			sieve[i] = true
		}
	}

	primes := make([]int, 0)
	primes = append(primes, 2)

	for i := 2; i < n; i++ {
		if sieve[i] {
			primes = append(primes, i)
			// Start from i^2, step by i, till n.
			// We mark i^2, i^2+i, i^2+2i, i^2+3i as composite.
			// We don't care for numbers less than i^2 as they'd already be marked,
			// and to prove a composite, we only need a single multiple > sqrt(i) for a number i.
			// 4, 6, 8, 10, 12...                [primes: 2, 3, 5, 7, 11, 13, 17, 19, 23]
			// 9, 12, 15, 18, 21...
			// 16, 20, 24, 28, 32...
			// 25, 30, 35, 40, 45...
			// 36, 42, 48, 54, 60...
			for j := i * i; j < len(sieve); j += i {
				sieve[j] = false
			}
		}
	}

	return primes
}

func main() {
	fmt.Println(sieveOfEratosthenes(100))
}
