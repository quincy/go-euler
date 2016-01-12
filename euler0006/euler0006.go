/*
Problem 6 from Project Euler.  http://projecteuler.net/problem=6

The sum of the squares of the first ten natural numbers is,

    1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

    (1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.
*/

package main

import "fmt"

func powers(min, max, power int64, nums chan int64) {
	for i := min; i <= max; i++ {
		nums <- i * i
	}

	close(nums)
}

func main() {
	var nums = make(chan int64)
	var min int64 = 1
	var max int64 = 100
	var squareOfSum int64
	var sumOfSquares int64

	// start calculating the squares in another thread
	go powers(min, max, 2, nums)

	// meanwhile calculate the square of the sum
	for i := min; i <= max; i++ {
		squareOfSum += i
	}

	squareOfSum *= squareOfSum

	// now find the sum of all of those squares from the other thread
	for i := range nums {
		sumOfSquares += i
	}

	// find the difference and print
	fmt.Println(squareOfSum - sumOfSquares)
}
