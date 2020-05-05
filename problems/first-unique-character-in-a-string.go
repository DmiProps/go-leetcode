package problems

import "fmt"

// Problem387 - 387. First Unique Character in a String
func Problem387() {

	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))

}

func firstUniqChar(s string) int {

	m := make(map[rune]int)
	for idx, ch := range s {
		if _, ok := m[ch]; ok {
			m[ch] = -1
		} else {
			m[ch] = idx
		}
	}
	ans := -1
	for _, idx := range m {
		if idx >= 0 && (ans < 0 || ans > idx) {
			ans = idx
		}
	}
	return ans

}
