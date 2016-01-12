/*
Problem 4 from Project Euler.  http://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"

// isPalindrome determines whether a given integer is a palindrome.
func isPalindrome(num int) bool {
	original := num
	reverse := 0

	for num > 0 {
		remainder := num % 10
		num /= 10
		reverse = reverse*10 + remainder
	}

	return original == reverse
}

// This works by making an educated guess that the two factors will be between
// 900..999.  So we don't bother searching any lower numbers.  The search also
// runs from in reverse across the range because the answer is more likely to
// be closer to 999.
func main() {
	for factorA := 999; factorA > 900; factorA-- {
		for factorB := 999; factorB > 900; factorB-- {
			product := factorA * factorB

			if isPalindrome(product) {
				fmt.Println(product)
				return
			}
		}
	}
}
