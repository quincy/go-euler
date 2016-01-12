/*
A perfect number is a number for which the sum of its proper divisors is
exactly equal to the number. For example, the sum of the proper divisors of 28
would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n
and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
number that can be written as the sum of two abundant numbers is 24. By
mathematical analysis, it can be shown that all integers greater than 28123 can
be written as the sum of two abundant numbers. However, this upper limit cannot
be reduced any further by analysis even though it is known that the greatest
number that cannot be expressed as the sum of two abundant numbers is less than
this limit.

Find the sum of all the positive integers which cannot be written as the sum of
two abundant numbers.

NOTE: This can likely be improved a lot.

$ time go run euler-problem23.go
4179871

real    0m2.197s
user    0m2.152s
sys     0m0.032s
*/

package main

import (
	"fmt"
)

// abundant determines if the sum of the proper divisors of n is greater than n.
func abundant(n int) (factors []int64, ok bool) {
	factors = properDivisors(int64(n))
	ok = int64(n) < sum(factors)
	return
}

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

func main() {
	// http://mathworld.wolfram.com/AbundantNumber.html
	// "Every number greater than 20161 can be expressed as a sum of two abundant numbers."
	limit := 20161
	anums := make(map[int][]int64)
	sumOfAnums := make(map[int]bool, limit)

	// Find all of the abundant numbers within the limit
	for i := 12; i <= limit; i++ {
		if factors, ok := abundant(i); ok {
			anums[i] = factors
		}
	}

	// Calculate the sum of each pair of abundant numbers.
	for k := range anums {
		for j := range anums {
			// k+j is an integer which CAN be written as the sum of two abundant numbers
			sumOfAnums[k+j] = true
		}
	}

	var sum int
	for i := 1; i <= limit; i++ {
		if _, ok := sumOfAnums[i]; !ok {
			sum += i
		}
	}

	fmt.Println(sum)
}
