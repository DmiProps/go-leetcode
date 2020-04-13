package problems

import "fmt"

// Problem525 - 525. Contiguous Array
func Problem525() {

	/*fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 0}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 0}))*/
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
	fmt.Println(findMaxLength([]int{1, 1, 1, 1, 1, 1, 1, 1, 0}))

}

func findMaxLength(nums []int) int {

	c := 0
	m := make(map[int]int)
	ans := 0

	for idx, n := range nums {
		if n == 0 {
			c++
		} else {
			c--
		}
		if c == 0 {
			ans = idx + 1
		} else if v, ok := m[c]; ok {
			if idx-v > ans {
				ans = idx - v
			}
		} else {
			m[c] = idx
		}
	}

	return ans

}
