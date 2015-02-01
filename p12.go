// p12: Highly divisible triangle number.
// 7th triangle number = 1 + 2 + 3 ... 7 = 28
// Factors of 28 = 1, 2, 4, 7, 14, 28
// 28 is the first triangle number to have over 5 divisors, find the number to have over 500.
// note that we count 1 and the number itself in the divisors!
// Sequence of triangle numbers goes as follows: 1, 3, 6, 10, 15, 21, 28...
// +build 12

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

func primeFactorize(n int, acc []int) []int {
	// Recursively accumulate all the prime factors
	var i int = 2
	if n/i == 0 {
		return acc
	}
	for i <= n {
		if n/i == 0 {
			acc = append(acc, i)
			return acc
		}
		if isPrime(i) && n%i == 0 {
			acc = append(acc, i)
			return primeFactorize(n/i, acc)
		} else {
			i += 1
		}
	}
	return acc
}

func countFactors(n int) int {
	// Count all factors using just the prime factors!
	factors := primeFactorize(n, make([]int, 0))
	midmap := make(map[int]int)
	for _, j := range factors {
		_, present := midmap[j]
		if !present {
			midmap[j] = 1
		} else {
			midmap[j] += 1
		}
	}

	var rv int = 1
	for _, v := range midmap {
		rv *= (v + 1)
	}
	return rv
}

func main() {
	found := false
	num := 1
	count := 2
	for !found {
		fmt.Println(num, ": ", countFactors(num))
		if countFactors(num) > 500 {
			fmt.Println("Found: ", num)
			found = true
		} else {
			// Get the next triangle number in the series.
			num += count
			count = count + 1
		}
	}
}
