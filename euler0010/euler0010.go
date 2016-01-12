/*
Problem 10 from Project Euler.  http://projecteuler.net/problem=10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

Answer: 142913828922
*/

package main

import "fmt"

func excludeMultiples(n int, sieve []bool) {
	for i := n; i < cap(sieve); i += n {
		sieve[i] = true
	}
}

func main() {
	var limit = 2000000
	var sum int64

	// sieve is zeroed, false indicates the value has not been excluded from
	// the sieve and therefore may be prime.  true indicates a known composite
	// number.
	var sieve = make([]bool, limit)

	// 0 and 1 are composite
	sieve[0] = true
	sieve[1] = true

	for i := 2; i < limit; {
		sum += int64(i)
		excludeMultiples(i, sieve)

		// find the next prime, which will always be the next false value in
		// the sieve
		for i < limit && sieve[i] {
			i++
		}
	}

	fmt.Println(sum)
}
