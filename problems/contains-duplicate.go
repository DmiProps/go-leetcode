package problems

import (
	"fmt"
	"sort"
)

// Problem217 - 217. Contains Duplicate
func Problem217() {
	test1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(test1))
	test2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(test2))
	test3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(test3))
}

func containsDuplicate(nums []int) bool {
	/*m := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}
	return false*/
	l := len(nums)
	if l < 2 {
		return false
	}
	sort.Ints(nums)
	for idx := 0; idx < l-1; idx++ {
		if nums[idx] == nums[idx+1] {
			return true
		}
	}
	return false

}
