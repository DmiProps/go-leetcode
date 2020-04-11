package problems

import "fmt"

// Problem11 - 11. Container With Most Water
func Problem11() {

	fmt.Println(maxArea([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

func maxArea(height []int) int {

	ans := 0
	for l, r := 0, len(height)-1; l < r; {
		n := min(height[l], height[r]) * (r - l)
		if n > ans {
			ans = n
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
