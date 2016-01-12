/*
The number, 1406357289, is a 0 to 9 pandigital number because it is made up of
each of the digits 0 to 9 in some order, but it also has a rather interesting
sub-string divisibility property.

Let d₁ be the 1st digit, d₂ be the 2nd digit, and so on. In this way, we note
the following:

    d₂d₃d₄  = 406 is divisible by 2
    d₃d₄d₅  = 063 is divisible by 3
    d₄d₅d₆  = 635 is divisible by 5
    d₅d₆d₇  = 357 is divisible by 7
    d₆d₇d₈  = 572 is divisible by 11
    d₇d₈d₉  = 728 is divisible by 13
    d₈d₉d₁₀ = 289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/

// FIXME This problem is not yet solved.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func intSliceToInt(ints []int) int {
	s := ""

	for _, v := range ints {
		s += strconv.Itoa(v)
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

type intPermuter struct {
	Seq []int
}

// First sets p.Seq equal to the first lexicographic permutation of p.Seq and
// returns p.Seq.
func (p *intPermuter) First() []int {
	sort.IntSlice(p.Seq).Sort()
	return p.Seq
}

// Next sets p.Seq equal to the next lexicographic permutation of p.Seq and
// returns p.Seq.
// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func (p *intPermuter) Next() ([]int, bool) {
	// Find the largest index m such that a[m] < a[m + 1]. If no such index
	// exists, the permutation is the last permutation.
	var m = -1
	for i, v := range p.Seq {
		if i == len(p.Seq)-1 {
			break
		} else if v < p.Seq[i+1] {
			m = i
		}
	}

	if m < 0 {
		return p.Seq, false
	}

	// Find the largest index n such that a[m] < a[n].
	var n = -1
	for i, v := range p.Seq {
		if p.Seq[m] < v {
			n = i
		}
	}

	// Swap the value of a[m] with that of a[n].
	p.Seq[m], p.Seq[n] = p.Seq[n], p.Seq[m]

	// Reverse the sequence from a[m + 1] up to and including the final element a[n].
	reverseOrder(p.Seq[m+1:])
	return p.Seq, true
}

// reverseOrder reverses the order of the elements in s in place.
func reverseOrder(s []int) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		tmp := s[i]
		s[i] = s[n-1-i]
		s[n-1-i] = tmp
	}
}

func excludeMultiples(n int, sieve []bool) {
	for i := n + n; i < len(sieve); i += n {
		sieve[i] = true
	}
}

func primeChecker(sieve []bool) func(int) bool {
	return func(n int) bool {
		return !sieve[n]
	}
}

// sieve is zeroed, false indicates the value has not been excluded from
// the sieve and therefore may be prime.  true indicates a known composite
// number.
var limit = 1000
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
	s := &intPermuter{Seq: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	s.First()

	ok := true
	for perm := s.First(); ok; perm, ok = s.Next() {
		if perm[0] == 0 {
			continue
		}
		isSolution := true
		for i := 1; i < 8; i++ {
			n, err := strconv.Atoi(string('0'+perm[i]) + string('0'+perm[i+1]) + string('0'+perm[i+2]))
			if err != nil {
				panic(err)
			}

			if !isPrime(n) {
				isSolution = false
			}
		}

		if isSolution {
			sum += intSliceToInt(perm)
		}
	}

	fmt.Println(sum)
}
