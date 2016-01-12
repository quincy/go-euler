/*
Starting in the top left corner of a 2×2 grid, and only being able to move to
the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
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

// choose calculates the combination of b unique sets from a things as from
// combinatorial mathematics.  a choose b == (a! / (b!)^2)
func choose(a, b *big.Int) (result *big.Int) {
	aFactorial := factorial(a)
	bFactorial := factorial(b)
	bFactSquared := big.NewInt(0)
	bFactSquared = bFactSquared.Exp(bFactorial, big.NewInt(2), &big.Int{})
	result = aFactorial.Div(aFactorial, bFactSquared)
	return
}

func main() {
	two := big.NewInt(2)
	rows := big.NewInt(20)

	rowsx2 := two.Mul(rows, two)

	// Calculate n choose r
	combinations := choose(rowsx2, rows)

	fmt.Printf("%s\n", combinations.String())
}
