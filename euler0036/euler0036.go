/*
The decimal number, 585 = 1001001001b (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in
base 10 and base 2.

(Please note that the palindromic number, in either base, may not include
leading zeros.)

Answer: 872187
*/

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	var mid int
	if len(s)%2 == 0 {
		mid = (len(s) / 2) - 1
	} else {
		mid = len(s) / 2
	}
	end := len(s) - 1

	for i := 0; i <= mid; i++ {
		if s[i] != s[end-i] {
			return false
		}
	}

	return true
}

func main() {
	sum := 0

	for i := 1; i < 1000000; i++ {
		base10 := strconv.FormatInt(int64(i), 10)
		base2 := strconv.FormatInt(int64(i), 2)

		if isPalindrome(base10) && isPalindrome(base2) {
			sum += i
		}
	}

	fmt.Println(sum)
}
