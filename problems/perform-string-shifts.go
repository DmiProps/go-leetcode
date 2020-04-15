package problems

import "fmt"

// ProblemX2 - Perform String Shifts
func ProblemX2() {

	fmt.Println(stringShift("abc", [][]int{{0, 1}, {1, 2}}))

	fmt.Println(stringShift("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}))

}

func stringShift(s string, shift [][]int) string {

	l := len(s)
	if l < 2 {
		return s
	}

	a := 0
	for _, i := range shift {
		if i[0] == 0 {
			a -= i[1]
		} else {
			a += i[1]
		}
	}
	a %= l
	if a == 0 {
		return s
	}
	if a < 0 {
		return s[-a:] + s[:-a]
	}
	return s[l-a:] + s[:l-a]

}
