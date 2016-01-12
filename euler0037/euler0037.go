/*
The number 3797 has an interesting property. Being prime itself, it is possible
to continuously remove digits from left to right, and remain prime at each
stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797,
379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to
right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
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

func truncatable(i int, isPrime func(int) bool) bool {
	s := strconv.Itoa(i)

	// Left to Right
	for i := 0; i < len(s); i++ {
		j, err := strconv.Atoi(s[i:])
		if err != nil {
			panic(err)
		}
		if !isPrime(j) {
			return false
		}
	}

	// Right to Left
	for i := len(s) - 1; i > 0; i-- {
		j, err := strconv.Atoi(s[:i])
		if err != nil {
			panic(err)
		}
		if !isPrime(j) {
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

	sum := 0
	count := 0
	for i := 11; i < len(sieve); i++ {
		if truncatable(i, isPrime) {
			sum += i
			count++
			if count == 11 {
				break
			}
		}
	}

	fmt.Println(sum)
}
