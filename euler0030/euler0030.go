/*
Surprisingly there are only three numbers that can be written as the sum of
fourth powers of their digits:

1634 = 1⁴ + 6⁴ + 3⁴ + 4⁴
8208 = 8⁴ + 2⁴ + 0⁴ + 8⁴
9474 = 9⁴ + 4⁴ + 7⁴ + 4⁴

As 1 = 1⁴ is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers
of their digits.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func sumOfFifthPower(n int) bool {
	s := strconv.Itoa(n)
	var sum int64

	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sum += int64(math.Pow(float64(i), 5))
	}

	if sum == int64(n) {
		return true
	}
	return false
}

func main() {
	var sum int
	var max = 200000

	for i := 10; i < max; i++ {
		if sumOfFifthPower(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
