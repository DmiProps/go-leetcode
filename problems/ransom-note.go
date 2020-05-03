package problems

import "fmt"

// Problem383 - 383. Ransom Note
func Problem383() {

	fmt.Println(canConstruct1("a", "b"))
	fmt.Println(canConstruct1("aa", "ab"))
	fmt.Println(canConstruct1("aa", "aab"))

}

func canConstruct1(ransomNote string, magazine string) bool {

	m := [26]byte{}
	for _, ch := range ransomNote {
		m[ch-'a']++
	}
	c := len(ransomNote)
	if c == 0 {
		return true
	}
	for _, ch := range magazine {
		if m[ch-'a'] > 0 {
			m[ch-'a']--
			c--
		}
		if c == 0 {
			return true
		}
	}
	return false

}
