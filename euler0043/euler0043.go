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
package euler0043

import (
	"sort"
	"strconv"
)

func IntSliceToInt(ints []int) int {
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

type IntPermuter struct {
	Seq []int
}

// First sets p.Seq equal to the first lexicographic permutation of p.Seq and
// returns p.Seq.
func (p *IntPermuter) First() []int {
	sort.IntSlice(p.Seq).Sort()
	return p.Seq
}

// Next sets p.Seq equal to the next lexicographic permutation of p.Seq and
// returns p.Seq.
// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func (p *IntPermuter) Next() ([]int, bool) {
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

func IsDivisible(values []int, target int) bool {
	a := 100 * values[0]
	b := 10 * values[1]
	c := values[2]

	return (a+b+c)%target == 0
}

var targets = []int{2, 3, 5, 7, 11, 13, 17}

func IsSubstringDivisible(input []int) bool {
	isSolution := true
	t := 0
	for i := 4; i <= len(input); i += 1 {
		target := targets[t]
		if !IsDivisible(input[i-3:i], target) {
			isSolution = false
			break
		}
		t++
	}

	return isSolution
}
