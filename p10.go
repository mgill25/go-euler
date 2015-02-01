// p10: Find the sum of all primes below 2 million
// +build 10

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	var i int
	var limit = int(math.Sqrt(float64(n)))
	for i = 2; i <= limit; i += 1 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 2
	num := 2000000

	for i := 3; i < num; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println("Sum is: ", sum)
}
