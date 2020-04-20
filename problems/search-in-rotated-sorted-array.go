package problems

import (
	"fmt"
)

// Problem33 - 33. Search in Rotated Sorted Array
func Problem33() {

	fmt.Println(search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)) // 4

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5)) // 1

	fmt.Println(search([]int{5, 1, 3}, 3)) // 2

	fmt.Println(search([]int{4, 5, 6, 7, 8, 0, 1, 2, 3}, 8)) // 4

	fmt.Println(search([]int{1, 3}, 3)) // 1

	fmt.Println(search([]int{5, 1, 3}, 5)) // 0

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1

	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 0)) // 0

	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 3)) // -1

}

func search(nums []int, target int) int {

	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, l-1
	for left <= right {
		if right-left < 2 {
			switch {
			case nums[left] == target:
				return left
			case nums[right] == target:
				return right
			default:
				return -1
			}
		}
		idx := left + (right-left)/2
		switch {
		case nums[idx] == target:
			return idx
		case nums[idx] < target:
			if target > nums[right] && nums[idx] < nums[left] {
				right = idx
			} else {
				left = idx
			}
		case nums[idx] > target:
			if target < nums[left] && nums[idx] > nums[right] {
				left = idx
			} else {
				right = idx
			}
		}
	}
	return -1

}
