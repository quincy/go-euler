/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of
their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.

Answer: 40730
*/

package main

import (
	"fmt"
	"strconv"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	total := 0
	fact := make(map[int]int)

	for i := 0; i < 10; i++ {
		fact[i] = factorial(i)
	}

	for i := 10; i < 50000; i++ {
		sum := 0
		s := strconv.Itoa(i)

		for _, v := range s {
			j, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			sum += fact[j]
		}

		if sum == i {
			total += i
		}
	}

	fmt.Println(total)
}
