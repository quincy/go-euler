/*
Problem 1 from Project Euler.  http://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
	"runtime"
)

// byThrees sums all of the multiples of 3 less than the limit.  The sum is
// written to the sum channel when complete.
func byThrees(sum chan int, limit int) {
	var mySum int

	for num := 3; num < limit; num += 3 {
		runtime.Gosched() // voluntarily yield the processor
		mySum += num
	}

	sum <- mySum // write the sum to the channel
}

// byFives sums all of the multiples of 5 less than the limit but exclude
// multiples of 3.  The sum is written to the channel sum when complete.
func byFives(sum chan int, limit int) {
	var mySum int

	for num := 5; num < limit; num += 5 {
		runtime.Gosched() // voluntarily yield the processor

		if num%3 != 0 {
			mySum += num
		}
	}

	sum <- mySum // write the sum to the channel
}

func main() {
	sum := make(chan int)
	limit := 1000

	go byThrees(sum, limit)
	go byFives(sum, limit)

	total := <-sum + <-sum

	fmt.Println(total)
}
