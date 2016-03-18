package main

import (
	"euler/euler0043"
	"fmt"
)

func main() {
	sum := 0
	s := &euler0043.IntPermuter{Seq: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	ok := true
	for perm := s.First(); ok; perm, ok = s.Next() {
		if euler0043.IsSubstringDivisible(perm) {
			sum += euler0043.IntSliceToInt(perm)
		}
	}

	fmt.Println(sum)
}
