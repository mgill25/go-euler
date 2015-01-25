// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder. What is the smallest positive number that is
// evenly divisible by all of the numbers from 1 to 20?
// Hint (basically the entire explanation) here: http://codereview.stackexchange.com/a/13091
// +build 5

package main

import (
	"fmt"
	"math"
)

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

func primeFactorize(n uint64, acc []uint64) []uint64 {
	// Recursively accumulate all the prime factors
	var i uint64 = 2
	if n/i == 0 {
		// If dividing by 2 yields 0, return what has been collected so far.
		return acc
	}

	// Start iterating...
	for i <= n {
		if n/i == 0 {
			// same as base case for current i. Append i to acc and return.
			acc = append(acc, i)
			return acc
		}

		if isPrime(i) && n%i == 0 {
			// if prime factor, append to acc and recurse.
			acc = append(acc, i)
			return primeFactorize(n/i, acc)
		} else {
			// otherwise, proceed with the loop
			i += 1
		}
	}
	// failed if here.
	return acc
}

func countPrimeFactors() map[uint64]map[uint64]uint64 {
	// Return dictionary mapping each number from 2 to 20 with its
	// prime count!
	// example: { '8': { 2: 3 } }
	dict := make(map[uint64]map[uint64]uint64)
	nums := []uint64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for _, i := range nums {
		factors := primeFactorize(i, make([]uint64, 0))
		// Count factors in array.
		midmap := make(map[uint64]uint64)
		for _, j := range factors {
			_, present := midmap[j]
			if !present {
				midmap[j] = 1
			} else {
				midmap[j] += 1
			}
		}
		dict[i] = midmap
	}
	return dict
}

func main() {
	var maxMap = make(map[uint64]uint64) // map for the final maximum exponents needed
	factorMap := countPrimeFactors()
	for _, factors := range factorMap {
		// fmt.Println(k, ": ", v)
		for num, exponent := range factors {
			// maxMap will contain the highest count of the exponents for
			// a number. So for example, if one number has 2 ** 3 and another has 2 ** 4
			// as factors, we only need the latter.
			_, present := maxMap[num]
			if !present {
				maxMap[num] = exponent
			} else {
				if exponent > maxMap[num] {
					maxMap[num] = exponent
				}
			}
		}
	}

	// Now we multiply all the max prime factors together
	var ans uint64 = 1
	for k, v := range maxMap {
		ans *= uint64(math.Pow(float64(k), float64(v)))
	}
	// >>> (2 ** 4) * 19 * 5 * (3**2) * 7 * 11 * 17 * 13
	// 232792560
	fmt.Println(ans)
}
