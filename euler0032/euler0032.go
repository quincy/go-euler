/*
We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can
be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only
include it once in your sum.
*/

/*
Strategy:
Factor each number from 1..987654321, then pass each set of multiplicand,
multiplier, and product to a function which checks to see if it is pandigital.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

// isPandigital tests the number in the string s to see if it is 1..n pandigital.
func isPandigital(s string, n int) bool {
	if len(s) != n {
		return false
	}

	digits := make([]bool, n+1)
	digits[0] = true

	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		if i < len(digits) {
			digits[i] = true
		}
	}

	for _, v := range digits {
		if !v {
			return false
		}
	}

	return true
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	m := make(map[int]bool)
	sum := 0

	// A 9-pandigital number is either a 2-digit · 3-digit number
	for a := 10; a < 99; a++ {
		for b := 100; b < 999; b++ {
			p := a * b
			if _, exists := m[p]; !exists {
				s := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(p)
				if isPandigital(s, 9) {
					m[p] = true
					sum += p
				}
			}
		}
	}

	// or a 1-digit · 4-digit number
	for a := 1; a < 9; a++ {
		for b := 1000; b < 9999; b++ {
			p := a * b
			if _, exists := m[p]; !exists {
				s := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(p)
				if isPandigital(s, 9) {
					m[p] = true
					sum += p
				}
			}
		}
	}

	fmt.Println(sum)
}
