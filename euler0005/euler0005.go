/*
Problem 5 from Project Euler.  http://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1
to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/

package main

import "fmt"

func main() {
	var result int64 = 2
	var i int64 = 3

	for ; i <= 20; i++ {
		result = lcm(result, i)
	}

	fmt.Println(result)
}

// gcf returns the Greatest Common Factor of two integers.
//
// It is implemented using Euclid's Algorithm.
func gcf(a, b int64) int64 {
	if a == b {
		return a
	} else if a > b {
		return gcf(a-b, b)
	}

	return gcf(b-a, a)
}

/*
lcm returns the Least Common Multiple two integers.
*/
func lcm(a, b int64) int64 {
	return a * b / gcf(a, b)
}
