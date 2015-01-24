// p2: Sum of Even fibonacci numbers below 4 million
// +build 2

package main

import (
	"fmt"
	"math"
)

func fib(p int) uint64 {
	// Calculate the n-th fibonacci number
	// using the Golden ratio formula.
	n := float64(p)
	phi := 1.618034
	rv := (math.Pow(phi, n) - math.Pow(1-phi, n)) / math.Sqrt(5)
	return uint64(rv)
}

func main() {
	var n uint64 = 4000000
	var sum uint64 = 0
	var num uint64 = 0

	// Iterate over even valued numbers below 4 million
	var i = 1
	for num < n {
		if num%2 == 0 {
			sum += num
		}
		// fmt.Printf("%d-th fibonacci number is: %d\n", i, num)
		num = fib(i) // do this *after* adding num to sum!
		i += 1
	}

	fmt.Printf("Sum is: %d\n", sum)
}
