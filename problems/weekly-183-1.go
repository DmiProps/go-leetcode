package problems

import (
	"fmt"
	"sort"
)

// Weekly183_1 - 5376. Minimum Subsequence in Non-Increasing Order
func Weekly183_1() {

	fmt.Println("Task 1...")

	test1 := []int{4, 3, 10, 9, 8}
	fmt.Println(minSubsequence(test1))

	test2 := []int{4, 4, 7, 6, 7}
	fmt.Println(minSubsequence(test2))

	test3 := []int{6}
	fmt.Println(minSubsequence(test3))

}

func minSubsequence(nums []int) []int {

	sort.Ints(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	ans := []int{}
	s := 0
	p := len(nums) - 1
	for s*2 <= sum {
		s += nums[p]
		ans = append(ans, nums[p])
		p--
	}
	return ans

}
