/*
The number, 197, is called a circular prime because all rotations of the
digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71,
73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"strconv"
)

func excludeMultiples(n int, sieve []bool) {
	for i := n + n; i < len(sieve); i += n {
		sieve[i] = true
	}
}

func isCircularPrime(n int, isPrime func(int) bool) bool {
	s := strconv.Itoa(n)

	for i := 1; i < len(s); i++ {
		s = s[1:] + string(s[0]) // rotate s
		p, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if !isPrime(p) {
			return false
		}
	}

	return true
}

func primeChecker(sieve []bool) func(int) bool {
	return func(n int) bool {
		return !sieve[n]
	}
}

// sieve is zeroed, false indicates the value has not been excluded from
// the sieve and therefore may be prime.  true indicates a known composite
// number.
var limit = 1000000
var sieve = make([]bool, limit)

func main() {
	var sum int

	// 0 and 1 are composite
	sieve[0] = true
	sieve[1] = true

	for i := 2; i < limit; {
		excludeMultiples(i, sieve)

		// find the next prime, which will always be the next false value in
		// the sieve
		i++
		for i < limit && sieve[i] {
			i++
		}
	}

	isPrime := primeChecker(sieve)

	for i := 2; i < len(sieve); i++ {
		if !sieve[i] {
			if isCircularPrime(i, isPrime) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
