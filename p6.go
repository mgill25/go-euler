// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
// +build 6

package main

import (
	"fmt"
	"math"
)

func sumOfSquares(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		square := int(math.Pow(float64(i), 2))
		sum += square
	}
	return sum
}

func squareOfSums(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func main() {
	num := 100
	diff := squareOfSums(num) - sumOfSquares(num)
	fmt.Println("Difference is: ", diff)
}
