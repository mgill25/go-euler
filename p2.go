// p2: Sum of Even fibonacci numbers below 4 million
// Hint: There is no need to start from 1 and 2 like the question
// can just start from 0 or 1.
// +build 2

package main

import (
	"fmt"
	"math"
)

func fib(n float64) uint64 {
	// Calculate the n-th fibonacci number
	// using the Golden ratio formula.
	phi := 1.618034
	// The formula is ((phi ^ n) - (1 - phi) ^ n) / sqrt(5)
	rv := (math.Pow(phi, n) - math.Pow(1-phi, n)) / math.Sqrt(5)
	return uint64(rv)
}

func main() {
	var n uint64 = 4000000
	var sum uint64 = 0
	var num uint64 = 0

	// Iterate over even valued numbers below 4 million
	var i int = 0 // start from 0
	for num < n {
		// We can skip the test for even because we are jumping by 3
		// and every 3rd fib is even.
		sum += num
		num = fib(float64(i)) // do this *after* adding num to sum!
		fmt.Printf("%d-th fibonacci number is: %d\n", i, num)
		i += 3 // jump by 3.
	}

	fmt.Printf("Sum is: %d\n", sum)
}
