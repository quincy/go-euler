/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into
alphabetical order. Then working out the alphabetical value for each name,
multiply this value by its alphabetical position in the list to obtain a name
score.

For example, when the list is sorted into alphabetical order, COLIN, which is
worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would
obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

// score calculates the name score for name.
func score(name string) (total int) {
	for _, c := range name {
		total += int(c-'A') + 1
	}
	return
}

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)

	var total int
	var names []string
	if names, err = reader.Read(); err != nil {
		panic(err)
	}
	sort.StringSlice(names).Sort()

	for k, v := range names {
		position := k + 1
		namescore := score(v)
		total += namescore * position
	}

	fmt.Println(total)
}
