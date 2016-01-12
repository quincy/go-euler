/*
The nth term of the sequence of triangle numbers is given by, t_n = ½n(n+1); so
the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value. For
example, the word value for SKY is 19 + 11 + 25 = 55 = t_10. If the word value
is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
containing nearly two-thousand common English words, how many are triangle
words?
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func wordScore(s string) (score int) {
	for _, v := range s {
		score += int(v-'A') + 1
	}

	return
}

func triangle(n int) int {
	return int(0.5 * float64(n) * (float64(n) + 1.0))
}

func main() {
	// read in the words
	f, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)

	var words []string
	if words, err = reader.Read(); err != nil {
		panic(err)
	}

	// Calculate the scores for each word
	m := make(map[string]int)

	highScore := 0
	for _, w := range words {
		m[w] = wordScore(w)
		if m[w] > highScore {
			highScore = m[w]
		}
	}

	// Calculate all triangle numbers up to the highest word score
	t := make(map[int]int)
	for i := 1; i <= highScore; i++ {
		t[triangle(i)] = i
	}

	// Remove any word whose score is not a triangle number
	for k, v := range m {
		if _, exists := t[v]; !exists {
			delete(m, k)
		}
	}

	fmt.Println(len(m))
}
