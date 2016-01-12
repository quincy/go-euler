/*
If p is the perimeter of a right angle triangle with integral length sides,
{a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
)

func solutions(p int) int {
	count := 0

	for i := 1; i < p-2; i++ {
		for j := 1; j < p-i-1; j++ {
			k := p - i - j
			if k > j { // Ensure we don't get the same answer twice.
				continue
			}
			if i*i == j*j+k*k {
				count++
			}
		}
	}

	return count
}

func main() {
	most := 0
	answer := 0

	for p := 1000; p > 0; p-- {
		sol := solutions(p)
		if sol > most {
			most = sol
			answer = p
		}
	}

	fmt.Println(answer)
}
