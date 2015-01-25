// p3: Largest Prime factor of the number 600851475143
// +build 3

// The prime factors of 13195 are 5, 7, 13 and 29.

package main

import (
	"fmt"
	"math"
)

// We'll write a function that returns an array/slice
// of all the prime factors of the input.
// Then we'll take the largest of those numbers as the answer.

func isPrime(n uint64) bool {
	if n == 0 || n == 1 {
		return false
	}
	var i uint64
	var limit = uint64(math.Sqrt(float64(n)))
	for i = 2; i <= limit; i += 1 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactorize(n uint64) []uint64 {

	var rv = make([]uint64, 0)
	// Start iterating primes
	var i uint64 = 2
	var limit = uint64(math.Sqrt(float64(n)))
	for i < limit {
		if isPrime(i) && n%i == 0 {
			rv = append(rv, i)
		}
		i += 1
	}
	return rv
}

func main() {
	// var num uint64 = 75
	var num uint64 = 600851475143
	var result = primeFactorize(num)
	fmt.Println("Result is: ", result)
	fmt.Println()
}
