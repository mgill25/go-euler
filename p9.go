// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a**2 + b**2 = c**2
// For example, 3**2 + 4**2 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
// +build 9

package main

import (
	"fmt"
	"math"
)

func iterateTriples() int {
	// The non-naive/non brute-force algorithm for this is a bit
	// complicated, using gcd or w/e. Check the PDF later...
	var result int
	found := false
	for m := 2; found != true; m++ {
		for n := 1; n < m; n++ {
			a := int(math.Pow(float64(m), 2)) - int(math.Pow(float64(n), 2))
			b := 2 * m * n
			c := int(math.Pow(float64(m), 2)) + int(math.Pow(float64(n), 2))
			if a+b+c == 1000 {
				found = true
				result = (a * b) * c
			}
		}
	}
	return result
}

func main() {
	// Euclid's formula to find Pythagorean triples:
	// Given 2 integers m and n, m > n
	// a = m^2 - n^2; b = 2mn; c = m^2 + n^2

	// The given example can be generated using the values m = 2, n = 1
	// We start from there and iterate until we find a triplet for which
	// the sum of a, b and c is 1000
	result := iterateTriples()
	fmt.Println("Result is: ", result)
}
