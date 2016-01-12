/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
letters. The use of "and" when writing out numbers is in compliance with
British usage.
*/

package main

import (
	"fmt"
)

var words = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

type wordNum int

func (w *wordNum) String() (result string) {
	num := int(*w)
	places := make([]int, 4)

	for i, j := 0, 1000; j >= 1; j /= 10 {
		places[i] = num / j
		num = num % j
		i++
	}

	if places[0] > 0 {
		result += words[places[0]] + " thousand "
	}
	if places[1] > 0 {
		result += words[places[1]] + " hundred "
		if places[2]+places[3] > 0 {
			result += "and "
		}
	}
	tensAndOnes := places[2]*10 + places[3]
	if tensAndOnes >= 10 && tensAndOnes < 20 {
		result += words[tensAndOnes]
		return
	} else if places[2] > 0 {
		result += words[places[2]*10]
	}
	if places[3] > 0 {
		if places[2] > 0 {
			result += "-"
		}
		result += words[places[3]]
	}

	return
}

func (w *wordNum) Count() (sum int) {
	str := w.String()
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			sum++
		}
	}

	return
}

func main() {
	var a = new(wordNum)
	var total int

	for i := 1; i <= 1000; i++ {
		*a = wordNum(i)
		total += a.Count()
	}

	fmt.Println(total)
}
