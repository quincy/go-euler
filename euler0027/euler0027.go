/*
Euler discovered the remarkable quadratic formula:

                                  n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values
n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible
by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80
primes for the consecutive values n = 0 to 79. The product of the coefficients,
−79 and 1601, is −126479.

Considering quadratics of the form:

                  n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n

                          e.g. |11| = 11 and |−4| = 4

Find the product of the coefficients, a and b, for the quadratic expression
that produces the maximum number of primes for consecutive values of n,
starting with n = 0.
*/

package main

import (
	"fmt"
)

func formula(a, b int) func(n int) int {
	return func(n int) int {
		return n*n + a*n + b
	}
}

func findConsecutivePrimes(a, b int) int {
	f := formula(a, b)
	n := 0
	_, ok := primes.IndexOf(f(n))
	for ok {
		n++
		_, ok = primes.IndexOf(f(n))
	}

	return n
}

// makePrimeGenerator returns a closure that outputs prime numbers.
//
// The prime numbers are generated naively.  Each new prime candidate is
// divided by all of the previously known primes.  If none of these divisions
// produces a whole number then the candidate is prime.  Otherwise it is
// composite.  The length of time it takes to generate each prime increases
// because of this.
func makePrimeGenerator() func() int {
	var primes = make([]int, 1000)
	var next int
	var current int
	var i = 0

	return func() int {
		// Primes list is seeded with 2 to begin
		if i == 0 {
			primes[i] = 2
			current = primes[i]
			next = 3
			i++
			return current
		}

		// increase the size of the primes slice if necessary
		if i == cap(primes)-1 {
			newSlice := make([]int, cap(primes)*2)
			copy(newSlice, primes)
			primes = newSlice
		}

		// test every odd number for primality until one is found, then return
		// it
		var isPrime = false
	CANDIDATE:
		for n := next; !isPrime; n += 2 {
			for p := 0; p < i; p++ {
				if n%primes[p] == 0 {
					continue CANDIDATE // n is composite, start testing next n
				}
			}

			// n must be prime so return it!
			isPrime = true
			primes[i] = n
			current = primes[i]
			next = n + 2
			i++
			break
		}

		return current
	}
}

type intMapList struct {
	values []int
	hash   map[int]int
}

func newIntMapList() *intMapList {
	var i = new(intMapList)
	i.Init()
	return i
}

func (l *intMapList) Init() {
	l.values = make([]int, 0)
	l.hash = make(map[int]int)
}

func (l *intMapList) Insert(i int) {
	l.hash[i] = len(l.values)
	l.values = append(l.values, i)
}

func (l *intMapList) IndexOf(i int) (index int, ok bool) {
	index, ok = l.hash[i]
	return
}

func (l *intMapList) Get(index int) (value int, ok bool) {
	if index < len(l.values) {
		value = l.values[index]
		ok = true
	}

	return
}

var primes = newIntMapList()
var p = makePrimeGenerator()

func main() {
	var most int
	var product int

	for i := 0; i < 1000; i++ {
		primes.Insert(p())
	}

	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			n := findConsecutivePrimes(a, b)
			if n > most {
				most = n
				product = a * b
			}
		}
	}

	fmt.Println(product)
}
