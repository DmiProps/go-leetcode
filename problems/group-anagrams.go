package problems

import (
	"fmt"
	"sort"
	"strings"
)

// ProblemDay6 - 49. Group Anagrams
func ProblemDay6() {

	test1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(test1))

}

func groupAnagrams(strs []string) [][]string {

	ans := [][]string{}
	tag := make(map[string]int)

	for _, s := range strs {

		s1 := sortSymbols(s)

		idx, ok := tag[s1]
		if ok {
			ans[idx] = append(ans[idx], s)
		} else {
			ans = append(ans, []string{s})
			tag[s1] = len(ans) - 1
		}

	}

	return ans

}

func sortSymbols(s string) string {

	s1 := strings.Split(s, "")
	sort.Strings(s1)

	s2 := make([]byte, len(s))
	for idx, ch := range s1 {
		s2[idx] = byte(ch[0])
	}
	return string(s2)

}
