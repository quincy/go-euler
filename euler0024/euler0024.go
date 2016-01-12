/*
A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
are listed numerically or alphabetically, we call it lexicographic order. The
lexicographic permutations of 0, 1 and 2 are:

                       012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5,
6, 7, 8 and 9?
*/

package main

import (
	"fmt"
	"sort"
)

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

const oneMillion = 1000000

func main() {
	var perm []int
	var ok bool
	s := &intPermuter{Seq: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	s.First()

	for i := 1; i < oneMillion; i++ {
		perm, ok = s.Next()
		if !ok {
			fmt.Println("Ran out of permutations.")
		}
	}

	fmt.Println(perm)
}
