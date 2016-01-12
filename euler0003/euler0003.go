/*
Problem 3 from Project Euler.  http://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	var num int64 = 600851475143

	var i int64 = 2

	for ; num > 1; i++ {
		for num%i == 0 {
			num /= i
		}
	}

	fmt.Println(i - 1)
}
