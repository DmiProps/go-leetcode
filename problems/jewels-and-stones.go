package problems

import (
	"fmt"
	"strings"
)

// ProblemX771 - 771. Jewels and Stones
func ProblemX771() {

	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0

}

func numJewelsInStones(J string, S string) int {

	ans := 0
	for _, ch := range S {
		if strings.Contains(J, string(ch)) {
			ans++
		}
	}
	return ans

}
