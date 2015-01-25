// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// +build 4

package main

import "fmt"

func isNumberPalindrome(num uint64) bool {
	// We can check this by re-creating the
	// number using modulo operator on the original
	// in reverse order.
	n := num
	var rev uint64 // reverse
	var rem uint64 // remainder
	for n > 0 {
		rem = n % 10
		rev = rev*10 + rem
		n = n / 10
	}
	if rev == num {
		return true
	}
	return false
}

func main() {
	var largest uint64 = 0
	// Iterate from 999 below.
	for i := 999; i >= 100; i-- {
		// To avoid repeated calculations, we can assume j <= i, so
		// that i * j happens once, and the second j * i never happens.
		for j := i; j >= 100; j-- {
			prod := uint64(i * j)
			if prod < largest {
				// We've found the largest, all prods from
				// now on will be lower, so stop calculating.
				break
			}
			if isNumberPalindrome(prod) && prod > largest {
				largest = prod
			}
		}
	}
	fmt.Println("Number is: ", largest)
}
