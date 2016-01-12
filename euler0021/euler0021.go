/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).

If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and
each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and
142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"
)

// properDivisors returns a []int64 of all of the properDivisors of num.
func properDivisors(num int64) []int64 {
	fmap := make(map[int64]bool)
	fmap[1] = true
	fmap[num] = true

	for i := int64(2); i < num; i++ {
		// If a number we are testing has already been found as a factor there
		// are no more factors to find.
		if _, exists := fmap[i]; exists {
			break
		}

		if num%i == 0 {
			fmap[i] = true
			fmap[num/i] = true
		}
	}

	factors := make([]int64, len(fmap)-1)
	i := 0
	for k := range fmap {
		if k < num {
			factors[i] = k
			i++
		}
	}

	return factors
}

// sum calculates the sum of a given slice.
func sum(nums []int64) (sum int64) {
	for _, v := range nums {
		sum += v
	}
	return
}

// amicable check if n has an amicable pair and returns its pair.
func amicable(n int64) (int64, bool) {
	m := sum(properDivisors(n))
	if n == sum(properDivisors(m)) && m != n {
		return m, true
	}

	return -1, false
}

func main() {
	var amicableNumbers = make(map[int64]int64)

	for i := int64(1); i < 10000; i++ {
		if _, exists := amicableNumbers[i]; !exists {
			if num, ok := amicable(i); ok {
				amicableNumbers[i] = num
				amicableNumbers[num] = i
			}
		}
	}

	var sum int64
	for k := range amicableNumbers {
		sum += k
	}

	fmt.Println(sum)
}
