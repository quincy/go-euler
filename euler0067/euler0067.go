/*
By starting at the top of the triangle below and moving to adjacent numbers on
the row below, the maximum total from top to bottom is 23.

                                       3
                                      7 4
                                     2 4 6
                                    8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and
'Save Link/Target As...'), a 15K text file containing a triangle with
one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible
to try every route to solve this problem, as there are 299 altogether! If you
could check one trillion (1012) routes every second it would take over twenty
billion years to check them all. There is an efficient algorithm to solve it.
;o)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	f, err := os.Open("triangle.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	var r int
	var triangle = make([][]int, 100)

	for scanner.Scan() {
		// read each row of ints from the file
		row := make([]int, r+1)
		rowText := scanner.Text()
		nums := strings.Split(rowText, " ")

		for i, v := range nums {
			// Convert the individual number strings into ints and add the
			// entire row to the triangle.
			if num, err := strconv.Atoi(v); err != nil {
				panic(err)
			} else {
				row[i] = num
			}

			// Now calculate the max possible sum so far.
			switch r {
			case 0:
			default:
				switch i {
				case 0:
					rightParent := triangle[r-1][i]
					row[i] = rightParent + row[i]
				case len(row) - 1:
					leftParent := triangle[r-1][i-1]
					row[i] = leftParent + row[i]
				default:
					leftParent := triangle[r-1][i-1]
					rightParent := triangle[r-1][i]
					row[i] = max(leftParent+row[i], rightParent+row[i])
				}
			}
		}

		// Add the new row to the triangle.
		triangle[r] = row
		r++
	}

	var maxVal int
	for _, v := range triangle[len(triangle)-1] {
		maxVal = max(maxVal, v)
	}

	fmt.Println(maxVal)
}
