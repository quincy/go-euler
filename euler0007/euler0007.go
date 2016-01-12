/*
Problem 7 from Project Euler.  http://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

What is the 10,001st prime number?
*/

package main

import "fmt"

// makePrimeGenerator returns a closure that outputs prime numbers.
//
// The prime numbers are generated naively.  Each new prime candidate is
// divided by all of the previously known primes.  If none of these divisions
// produces a whole number then the candidate is prime.  Otherwise it is
// composite.  The length of time it takes to generate each prime increases
// because of this.
func makePrimeGenerator() func() int {
	var primes = make([]int, 1000)
	var next int
	var current int
	var i = 0

	return func() int {
		// Primes list is seeded with 2 to begin
		if i == 0 {
			primes[i] = 2
			current = primes[i]
			next = 3
			i++
			return current
		}

		// increase the size of the primes slice if necessary
		if i == cap(primes)-1 {
			newSlice := make([]int, cap(primes)*2)
			copy(newSlice, primes)
			primes = newSlice
		}

		// test every odd number for primality until one is found, then return
		// it
		var isPrime = false
	CANDIDATE:
		for n := next; !isPrime; n += 2 {
			for p := 0; p < i; p++ {
				if n%primes[p] == 0 {
					continue CANDIDATE // n is composite, start testing next n
				}
			}

			// n must be prime so return it!
			isPrime = true
			primes[i] = n
			current = primes[i]
			next = n + 2
			i++
			break
		}

		return current
	}
}

func main() {
	var prime int
	nextPrime := makePrimeGenerator()

	var limit = 10001

	for i := 0; i < limit; i++ {
		prime = nextPrime()
	}

	fmt.Println(prime)
}
