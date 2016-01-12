/*
2¹⁵ = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2¹⁰⁰⁰?
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := big.NewInt(2)
	num.Exp(num, big.NewInt(1000), &big.Int{})

	var total int64
	for _, v := range num.String() {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		total += int64(i)
	}

	fmt.Println(total)
}
