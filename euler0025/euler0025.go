/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F₁ = 1 and F₂ = 1.
Hence the first 12 terms will be:

F₁  = 1
F₂  = 1
F₃  = 2
F₄  = 3
F₅  = 5
F₆  = 8
F₇  = 13
F₈  = 21
F₉  = 34
F₁₀ = 55
F₁₁ = 89
F₁₂ = 144

The 12th term, F₁₂, is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import (
	"fmt"
	"math/big"
)

// fib returns a function that returns successive Fibonacci numbers.
// Source: golang.org
func fib() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	return func() *big.Int {
		c := big.NewInt(0)
		a, b = b, c.Add(a, b)
		return a
	}
}

func main() {
	f := fib()

	for i := 1; ; i++ {
		next := f()
		if len(next.String()) == 1000 {
			fmt.Println(i)
			break
		}
	}
}
