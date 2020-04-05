package problems

import (
	"fmt"
)

// Biweekly23_2 - 1400. Construct K Palindrome Strings
func Biweekly23_2() {

	fmt.Println(canConstruct("annabelle", 2))
	fmt.Println(canConstruct("leetcode", 3))
	fmt.Println(canConstruct("true", 4))
	fmt.Println(canConstruct("yzyzyzyzyzyzyzy", 2))
	fmt.Println(canConstruct("cr", 7))

}

func canConstruct(s string, k int) bool {

	if k > len(s) {
		return false
	}
	if k == len(s) {
		return true
	}

	m := make(map[byte]int)
	cnt := 0
	for i := 0; i < len(s); i++ {
		c := m[s[i]] + 1
		if c%2 == 1 {
			cnt++
		} else {
			cnt--
		}
		m[s[i]] = c
	}
	if cnt > k {
		return false
	}
	return true

}
