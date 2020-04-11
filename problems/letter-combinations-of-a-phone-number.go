package problems

import "fmt"

// Problem17 - 17. Letter Combinations of a Phone Number
func Problem17() {

	ans := letterCombinations("23")

	for _, s := range ans {
		fmt.Print(s, ", ")
	}

}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	m := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	c := 1
	l := len(digits)
	for _, s := range digits {
		c *= len(m[s-'2'])
	}
	ans := make([]string, c)
	str := make([]byte, l)
	for i := 0; i < c; i++ {
		d := i
		for j := 0; j < l; j++ {
			s := len(m[digits[j]-'2'])
			str[j] = m[digits[j]-'2'][d%s]
			d /= s

		}
		ans[i] = string(str)
	}
	return ans

}
