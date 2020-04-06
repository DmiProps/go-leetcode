package problems

import "fmt"

// Problem26 - 26. Remove Duplicates from Sorted Array
func Problem26() {

	fmt.Println(removeDuplicates([]int{1, 1, 2}))

	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

}

func removeDuplicates(nums []int) int {

	len := 0
	for idx, n := range nums {
		if idx == 0 {
			len = 1
			continue
		}
		if idx >= len && n > nums[len-1] {
			nums[len] = n
			len++
		}

	}
	return len

}
