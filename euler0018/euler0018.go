/*
By starting at the top of the triangle below and moving to adjacent numbers on
the row below, the maximum total from top to bottom is 23.

                                       3
                                      7 4
                                     2 4 6
                                    8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                                       75
                                     95 64
                                    17 47 82
                                  18 35 87 10
                                 20 04 82 47 65
                               19 01 23 75 03 34
                              88 02 77 73 07 63 67
                            99 65 04 28 06 16 70 92
                           41 41 26 56 83 40 80 70 33
                         41 48 72 33 47 32 37 16 94 29
                        53 71 44 65 25 43 91 52 97 51 14
                      70 11 33 28 77 73 17 78 39 68 17 57
                     91 71 52 38 17 14 91 43 58 50 27 29 48
                   63 66 04 68 89 53 67 30 73 16 69 87 40 31
                  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by
trying every route. However, Problem 67, is the same challenge with a triangle
containing one-hundred rows; it cannot be solved by brute force, and requires a
clever method! ;o)
*/

package main

import (
	"fmt"
)

var triangle = [][]int{
	[]int{75},
	[]int{95, 64},
	[]int{17, 47, 82},
	[]int{18, 35, 87, 10},
	[]int{20, 04, 82, 47, 65},
	[]int{19, 01, 23, 75, 03, 34},
	[]int{88, 02, 77, 73, 07, 63, 67},
	[]int{99, 65, 04, 28, 06, 16, 70, 92},
	[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
	[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	[]int{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	[]int{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	for r, row := range triangle {
		for v, value := range row {
			switch r {
			case 0:
			default:
				switch v {
				case 0:
					rightParent := triangle[r-1][v]
					row[v] = rightParent + value
				case len(row) - 1:
					leftParent := triangle[r-1][v-1]
					row[v] = leftParent + value
				default:
					leftParent := triangle[r-1][v-1]
					rightParent := triangle[r-1][v]
					row[v] = max(leftParent+value, rightParent+value)
				}
			}
		}
	}

	var maxVal int
	for _, v := range triangle[len(triangle)-1] {
		maxVal = max(maxVal, v)
	}

	fmt.Println(maxVal)
}
