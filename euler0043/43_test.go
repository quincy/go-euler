package euler0043

import "testing"

func IsSubstringDivisible_test(t *testing.T) {
	tests := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}, true},
		{[]int{1, 4, 0, 3, 6, 5, 7, 2, 8, 9}, false},
	}

	for _, test := range tests {
		got := IsSubstringDivisible(test.in)
		if got != test.want {
			t.Errorf("IsSubstringDivisible(%v) want %d got %d", test.in, test.want, got)
		}
	}
}

func IntSliceToInt_test(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}, 1406357289},
		{[]int{1, 4, 0, 3, 6, 5, 7, 2, 8, 9}, 1403657289},
	}

	for _, test := range tests {
		got := IntSliceToInt(test.in)
		if got != test.want {
			t.Errorf("IntSliceToInt(%v) want %d got %d", test.in, test.want, got)
		}
	}
}

func IsDivisible_test(t *testing.T) {
	tests := []struct {
		in     []int
		target int
		want   bool
	}{
		{[]int{4, 0, 6}, 2, 406%2 == 0},
		{[]int{4, 0, 6}, 3, 406%3 == 0},
		{[]int{4, 0, 6}, 5, 406%5 == 0},
		{[]int{4, 0, 6}, 7, 406%7 == 0},
		{[]int{4, 0, 6}, 11, 406%11 == 0},
		{[]int{4, 0, 6}, 13, 406%13 == 0},
		{[]int{4, 0, 6}, 17, 406%17 == 0},

		{[]int{0, 4, 6}, 2, 46%2 == 0},
		{[]int{0, 4, 6}, 3, 46%3 == 0},
		{[]int{0, 4, 6}, 5, 46%5 == 0},
		{[]int{0, 4, 6}, 7, 46%7 == 0},
		{[]int{0, 4, 6}, 11, 46%11 == 0},
		{[]int{0, 4, 6}, 13, 46%13 == 0},
		{[]int{0, 4, 6}, 17, 46%17 == 0},
	}

	for _, test := range tests {
		got := IsDivisible(test.in, test.target)
		if got != test.want {
			t.Errorf("%v %% %d want %b got %b", test.in, test.target, test.want, got)
		}
	}
}
