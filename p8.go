// Problem 8
// Answer using clever string substitution: 23514624000
// Great explanation here: http://codereview.stackexchange.com/a/59478
// 2541865828329
// +build 8
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getProduct(arr []int) int {
	ans := 1
	for i := 0; i < len(arr); i++ {
		ans *= arr[i]
	}
	return ans
}

func main() {
	s := `73167176531330624919225119674426574742355349194934
		96983520312774506326239578318016984801869478851843
		85861560789112949495459501737958331952853208805511
		12540698747158523863050715693290963295227443043557
		66896648950445244523161731856403098711121722383113
		62229893423380308135336276614282806444486645238749
		30358907296290491560440772390713810515859307960866
		70172427121883998797908792274921901699720888093776
		65727333001053367881220235421809751254540594752243
		52584907711670556013604839586446706324415722155397
		53697817977846174064955149290862569321978468622482
		83972241375657056057490261407972968652414535100474
		82166370484403199890008895243450658541227588666881
		16427171479924442928230863465674813919123162824586
		17866458359124566529476545682848912883142607690042
		24219022671055626321111109370544217506941658960408
		07198403850962455444362981230987879927244284909188
		84580156166097919133875499200524063689912560717606
		05886116467109405077541002256983155200055935729725
		71636269561882670428252483600823257530420752963450`

	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	// Brute-forcing. Iterate through set of 13 digit numbers and keep track
	// of maximum product seen so far.
	maxProduct := 0
	totalCount := 13

	circularBuffer := make([]int, totalCount)
	// Circular or Ring buffer: http://en.wikipedia.org/wiki/Circular_buffer
	// Nice property: Once filled, starts over-writing old data. Using fill-count
	// instead of end-pointer.

	for index, bufferIndex := 0, 0; index < len(s); index++ {
		// if bufferIndex == totalCount {
		// 	bufferIndex = 0 // can't access more than totalCount index in buffer.
		// }

		bufferIndex = bufferIndex % totalCount // can also toggle back to beginning with modulo.

		num, err := strconv.Atoi(string(s[index]))
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		circularBuffer[bufferIndex] = num // add the next number in buffer

		if index < totalCount-1 {
			// circular buffer not fully initialized yet.
			continue
		}

		// all numbers are in buffer here...
		prod := getProduct(circularBuffer)
		if prod > maxProduct {
			maxProduct = prod
		}

		bufferIndex += 1 // can't update in for statement. :/
	}
	fmt.Println("maxProduct is: ", maxProduct)
}
