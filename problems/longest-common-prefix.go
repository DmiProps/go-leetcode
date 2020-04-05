package problems

import "fmt"

// Problem14 - 14. Longest Common Prefix
func Problem14() {

	test1 := []string{"flower", "flow", "flight"}
	fmt.Println("Prefix:", longestCommonPrefix(test1))

	test2 := []string{"dog", "racecar", "car"}
	fmt.Println("Prefix:", longestCommonPrefix(test2))

}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	pref := strs[0]
	l := len(pref)

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(pref) {
			l = len(strs[i])
			pref = string(pref[:l])
		}
		if l == 0 {
			return ""
		}
		for {
			if pref[:l] == strs[i][:l] {
				if l < len(pref) {
					pref = string(pref[:l])
				}
				break
			} else {
				l--
				if l == 0 {
					return ""
				}
			}
		}
	}

	return pref

}
