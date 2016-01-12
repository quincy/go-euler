/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in
attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is
correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than
one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms,
find the value of the denominator.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isReverse(a, b string) bool {
	j := len(b) - 1
	for k, v := range a {
		if v != rune(b[j-k]) {
			return false
		}
	}

	return true
}

func hasCommonDigit(a, b string) bool {
	for _, v := range a {
		if strings.ContainsRune(b, v) {
			return true
		}
	}

	return false
}

func getCommonDigit(a, b int) (int, bool) {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)

	for _, v := range sa {
		if strings.ContainsRune(sb, v) {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			return i, true
		}
	}

	return 0, false
}

func isCandidate(a, b int) bool {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)

	return hasCommonDigit(sa, sb) && !isReverse(sa, sb)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % t
		a = t
	}
	return a
}

func reduce(n, d int) (num, den int) {
	divisor := gcd(n, d)
	num = n / divisor
	den = d / divisor
	return
}

func stripDigit(a, b, d int) (int, int) {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)
	sd := strconv.Itoa(d)
	sa = strings.Trim(sa, sd)
	sb = strings.Trim(sb, sd)

	var err error
	a, err = strconv.Atoi(sa)
	if err != nil {
		panic(err)
	}
	b, err = strconv.Atoi(sb)
	if err != nil {
		panic(err)
	}

	return a, b
}

func naiveReduce(n, d int) (num, den int) {
	if i, ok := getCommonDigit(n, d); ok {
		num, den = stripDigit(n, d, i)
		return reduce(num, den)
	}
	return n, d
}

func fracMult(n, d, p, q int) (int, int) {
	num := n * p
	den := d * q
	return reduce(num, den)
}

func main() {
	rn := 1
	rd := 1

	for n := 12; n < 100; n++ {
		if n%10 == 0 || n%11 == 0 {
			continue
		}
		for d := 12; d < 100; d++ {
			if d%10 == 0 || d%11 == 0 {
				continue
			}
			if n < d {
				p, q := naiveReduce(n, d)
				v, w := reduce(n, d)
				//fmt.Printf("Candidate        %d/%d\n", n, d)
				//fmt.Printf("Actual reduction %d/%d\n", v, w)
				//fmt.Printf("Naive  reduction %d/%d\n", p, q)

				if p == v && q == w && p != n && q != d {
					//fmt.Printf("Accepted         %d/%d\n", n, d)
					rn, rd = fracMult(rn, rd, n, d)
				}
			}
		}
	}

	fmt.Printf("%d/%d\n", rn, rd)
}
