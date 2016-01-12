/*
A unit fraction contains 1 in the numerator. The decimal representation of the
unit fractions with denominators 2 to 10 are given:

1/2  = 0.5
1/3  = 0.(3)
1/4  = 0.25
1/5  = 0.2
1/6  = 0.1(6)
1/7  = 0.(142857)
1/8  = 0.125
1/9  = 0.(1)
1/10 = 0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be
seen that 1/7 has a 6-digit recurring cycle.

Find the value of d<1000 for which 1/d contains the longest recurring cycle in
its decimal fraction part.
*/

package main

import (
	"fmt"
	"strconv"
)

func sequenceLength(n, d int) (length int, digits string) {
	rems := make(map[int]int)
	i := 0
	q := n / d
	n = n % d
	rems[n] = i
	digits += strconv.Itoa(q) + "."

	for i = 1; i < d; i++ {
		n *= 10
		q = n / d
		n %= d

		digits += strconv.Itoa(q)

		if r, ok := rems[n]; ok {
			length = i - r
			break
		}
		rems[n] = i
	}

	return
}

func main() {
	var longest int
	var length int
	var value string

	for d := 1; d < 1000; d++ {
		seq, digits := sequenceLength(1, d)
		if seq > length {
			longest = d
			length = seq
			value = fmt.Sprintf("%s(%s)", digits[:len(digits)-length], digits[len(digits)-length:])
		}
	}
	fmt.Printf("1/%d has the longest cycle of length %d.  1/%d=%s\n", longest, length, longest, value)
}
