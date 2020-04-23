package problems

import "fmt"

// Problem201 - 201. Bitwise AND of Numbers Range
func Problem201() {

	fmt.Println(rangeBitwiseAnd(3, 4)) // 0
	fmt.Println(rangeBitwiseAnd(5, 7)) // 4
	fmt.Println(rangeBitwiseAnd(0, 1)) // 0

}

func rangeBitwiseAnd(m int, n int) int {

	ans := m

	for c, t := 1, 0; c <= n; c <<= 1 {
		if m&c != 0 {
			if (m|t)+1 <= n {
				ans -= c
			}
		}
		t = t<<1 + 1
	}

	return ans

}
