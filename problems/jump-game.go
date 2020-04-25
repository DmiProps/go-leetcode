package problems

import "fmt"

// Problem55 - 55. Jump Game
func Problem55() {

	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0}))             // true

}

func canJump(nums []int) bool {

	s := 0
	for idx, n := range nums {
		if idx+n > s {
			s = idx + n
		}
		if idx >= s && idx < len(nums)-1 {
			return false
		}
	}
	return true

}
