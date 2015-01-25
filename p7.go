// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
// that the 6th prime is 13.
// What is the 10001st prime number?
// +build 7

package main

import (
	"fmt"
	"math"
)

func isPrime(n uint64) bool {
	if n == 0 || n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	var i uint64
	var limit = uint64(math.Sqrt(float64(n)))
	for i = 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nthPrime(n uint64) uint64 {
	if n == 1 {
		return 2 // handle the only even prime
	}
	var i uint64 = 3
	var count uint64 = 1
	for {
		if isPrime(i) {
			count += 1
			// fmt.Println(count, "th prime is: ", i)
		}
		if count == n {
			return i
		}
		i += 2 // jump by 2 because looking at odd numbers only.
	}
}

func main() {
	var num uint64 = 10001
	fmt.Println("The ", num, "st prime number is: ", nthPrime(num))
}
