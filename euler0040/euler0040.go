/*
An irrational decimal fraction is created by concatenating the positive
integers:

                     0.123456789101112131415161718192021...
                                  ^

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the
following expression.

d_1 × d_10 × d_100 × d_1000 × d_10000 × d_100000 × d_1000000
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	lim := 100000
	s := " "

	for i := 1; len(s) <= lim; i++ {
		s = s + strconv.Itoa(i)
	}

	p := 1
	for i := 1; i <= lim; i *= 10 {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		p *= n
	}
	fmt.Println(p)
}
