/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even) n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1 It can be seen that this sequence
(starting at 13 and finishing at 1) contains 10 terms. Although it has not been
proved yet (Collatz Problem), it is thought that all starting numbers finish at
1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import "fmt"

// collatz calculates the collatz sequence for num and returns it.
func collatz(num int) []int {
	var seq = make([]int, 100)

	seq[0] = num

	var i = 1
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}

		// Re-allocate the slice if necessary.
		if i >= cap(seq) {
			newSeq := make([]int, cap(seq)*2)
			copy(newSeq, seq)
			seq = newSeq
		}

		seq[i] = num
		i++
	}

	return seq[:i]
}

func main() {
	longest := 0
	start := 0
	for i := 1; i < 1000000; i++ {
		seq := collatz(i)
		if len(seq) > longest {
			longest = len(seq)
			start = i
		}
	}

	fmt.Println(start)
}
