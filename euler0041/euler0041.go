/*
We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is
also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import (
	"fmt"
	"strconv"
)

// isPandigital tests the number in the string s to see if it is 1..n pandigital.
func isPandigital(s string, n int) bool {
	if len(s) != n {
		return false
	}

	digits := make([]bool, n+1)
	digits[0] = true

	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		if i < len(digits) {
			digits[i] = true
		}
	}

	for _, v := range digits {
		if !v {
			return false
		}
	}

	return true
}

func excludeMultiples(n int, sieve []bool) {
	for i := n + n; i < len(sieve); i += n {
		sieve[i] = true
	}
}

// sieve is zeroed, false indicates the value has not been excluded from
// the sieve and therefore may be prime.  true indicates a known composite
// number.
var limit = 8000000
var sieve = make([]bool, limit)

func main() {
	answer := 0

	// 0 and 1 are composite
	sieve[0] = true
	sieve[1] = true

	for i := 2; i < limit; {
		s := strconv.Itoa(i)
		if isPandigital(s, len(s)) {
			answer = i
		}

		excludeMultiples(i, sieve)

		// find the next prime, which will always be the next false value in
		// the sieve
		i++
		for i < limit && sieve[i] {
			i++
		}
	}

	fmt.Println(answer)
}
