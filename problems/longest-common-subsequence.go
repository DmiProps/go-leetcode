package problems

import "fmt"

// Problem1143 - 1143. Longest Common Subsequence
func Problem1143() {

	fmt.Println(longestCommonSubsequence("abcde", "ace")) // 3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   // 3
	fmt.Println(longestCommonSubsequence("abc", "def"))   // 0

}

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	m := make([][]int, l1)
	ans := 0
	for i := 0; i < l1; i++ {
		m[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					m[i][j] = m[i-1][j-1] + 1
				} else {
					m[i][j] = 1
				}
			} else {
				a := 0
				if i > 0 {
					a = m[i-1][j]
				}
				if j > 0 && m[i][j-1] > a {
					a = m[i][j-1]
				}
				m[i][j] = a
			}
			if m[i][j] > ans {
				ans = m[i][j]
			}
		}
	}
	return ans
}
