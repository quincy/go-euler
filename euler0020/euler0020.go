/*
n! means × n × (n × 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"math/big"
)

// factorial calculates the factorial of num and returns it.
func factorial(num *big.Int) (result *big.Int) {
	one := big.NewInt(1)
	result = big.NewInt(1)
	n := big.NewInt(num.Int64())

	for n.Cmp(&big.Int{}) == 1 {
		result.Mul(result, n)
		n.Sub(n, one)
	}
	return
}

func main() {
	var sum int

	result := factorial(big.NewInt(100))
	for _, v := range result.String() {
		sum += int(v - 48)
	}

	fmt.Println(sum)
}
