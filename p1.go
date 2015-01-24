// p1: Sum of all multiples of 3 or 5 below 1000

package main

import "fmt"

func sumDivisibleBy(n int) int {
	target := 999
	p := target / n
	rv := n * (p * (p + 1)) / 2
	return rv
}

func main() {
	// sum := 0
	// for i := 0; i < target; i++ {
	// 	if i%3 == 0 || i%5 == 0 {
	// 		sum += i
	// 	}
	// }
	sum := sumDivisibleBy(3) + sumDivisibleBy(5) - sumDivisibleBy(15)
	fmt.Printf("Sum is %d\n", sum)
}
