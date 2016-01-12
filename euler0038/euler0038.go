/*
Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will
call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and
5, giving the pandigital, 918273645, which is the concatenated product of 9 and
(1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
	"strconv"
)

// isPandigital tests the number in the string s to see if it is 1..n pandigital.
func isPandigital(s string, n int) bool {
	if len(s) != n {
		return false
	}

	digits := make([]bool, n+1)
	digits[0] = true

	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		if i < len(digits) {
			digits[i] = true
		}
	}

	for _, v := range digits {
		if !v {
			return false
		}
	}

	return true
}

func main() {
	limit := 10000
	largest := 0

	for i := 1; i < limit; i++ {
		s := strconv.Itoa(i)

		for j := 2; len(s) < 9; j++ {
			s = s + strconv.Itoa(i*j)
		}

		if isPandigital(s, 9) {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			if n > largest {
				largest = n
			}
		}
	}

	fmt.Println(largest)
}
