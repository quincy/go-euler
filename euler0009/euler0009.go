/*
Problem 9 from Project Euler.  http://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

    a² + b² = c²

For example, 3² + 4² = 9 + 16 = 25 = 5².

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

Answer: 31875000

According to Wikipedia, every Pythagorean triplet is expressed by

    a = m² - n², b = 2mn, c = m² + n²

So a brute force approach is to increment m and n one at a time and find a, b,
and c.
*/

package main

import "fmt"

// Calculate a Pythagorean triplet.
func PythagoreanTriplet(m, n int) (a, b, c int) {
	a = m*m - n*n
	b = 2 * m * n
	c = m*m + n*n
	return
}

func main() {
	for n := 0; n < 500; n++ {
		for m := 0; m < 500; m++ {
			a, b, c := PythagoreanTriplet(m, n)

			if (a + b + c) == 1000 {
				fmt.Printf("%v² + %v² = %v² and %v + %v + %v = 1000\n", a, b, c, a, b, c)
				fmt.Println(a * b * c)
				return
			}
		}
	}

	fmt.Println("No Pythagorean triplets were found which sum to 1000.")
}
