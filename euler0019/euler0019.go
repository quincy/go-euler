/*
You are given the following information, but you may prefer to do some research
for yourself.

 * 1 Jan 1900 was a Monday.
 * Thirty days has September,
   April, June and November.
   All the rest have thirty-one,
   Saving February alone,
   Which has twenty-eight, rain or shine.
   And on leap years, twenty-nine.
 * A leap year occurs on any year evenly divisible by 4, but not on a century
   unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1
Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
)

const (
	january = iota
	february
	march
	april
	may
	june
	july
	august
	september
	october
	november
	december
	leapFeb
)

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

var numDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 29}

// isLeapYear returns whether or not a given year is a leap year.
func isLeapYear(y int) bool {
	if y%4 == 0 && y%400 != 0 {
		return true
	}

	return false
}

// weekday returns the zero-based weekday number given an offset from a
// starting weekday number.
func weekday(current, offset int) int {
	return (current + offset) % 7
}

func main() {
	total := 0
	wday := monday

	for year := 1900; year < 2000; year++ {
		for month := january; month <= december; month++ {
			offset := numDays[month]

			if isLeapYear(year) && month == february {
				offset = numDays[leapFeb]
			}

			wday = weekday(wday, offset)

			if wday == sunday {
				total++
			}
		}
	}

	fmt.Println(total)
}
