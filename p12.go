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
	"time"
)

func Primes(n int) []int {
	// Generate Prime numbers upto n using Sieve of Eratosthenes.
	sieve := make([]bool, n)
	for i := 0; i < n; i++ {
		sieve[i] = true
	}

	// now cross off the non-primes.
	primes := make([]int, 0)
	for i := 2; i < n; i++ {
		if sieve[i] {
			primes = append(primes, i)
			for j := i * i; j < n; j += i {
				sieve[j] = false
			}
		}
	}
	return primes
}

func countDivisors(n int, primes []int, divisors int) int {
	// helper function
	// keep reducing the number until it becomes 1
	// Iterate over each prime number and count its exponent as it is able
	// to divide the remaining quotient. Add one to each prime's exponent
	// and multiply them at the end.
	// Ref: http://www.wikihow.com/Determine-the-Number-of-Divisors-of-an-Integer

	k := 0       // prime index
	for n != 1 { // keep reducing the number until it becomes 1
		count := 1
		for n%primes[k] == 0 {
			n /= primes[k]
			count++
		}
		divisors *= count
		k++
	}
	return divisors
}

func countFactors(n int, primes []int) int {
	// Every prime number only has 2 factors: Itself and 1.
	for i := 0; i < len(primes); i++ {
		if primes[i] == n {
			return 2
		}
	}

	// if non-prime, we use formula!
	return countDivisors(n, primes, 1)
}

func triangular(n int) int {
	return n * (n + 1) / 2
}

func main() {
	t0 := time.Now()

	num := 1

	// generate primes via Sieve of Eratosthenes
	primeLen := 100000
	primes := Primes(primeLen)

	triangular_number := triangular(num)
	num_divisors := countFactors(triangular_number, primes)

	for num_divisors <= 500 {
		// fmt.Println(triangular_num, ": ", num_divisors)
		num += 1
		triangular_number = triangular(num)
		num_divisors = countFactors(triangular_number, primes)
	}

	fmt.Println("Found: ", triangular_number, " with ", num_divisors, " divisors")
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

	// fmt.Println("Divisors of 16: ", countDivisors(16, primes, 1)) 	// 2^4 * 1^1
}
