package problems

import (
	"fmt"
	"sort"
)

// Problem18 - 18. 4Sum
func Problem18() {

	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))

}

func fourSum(nums []int, target int) [][]int {

	ans := [][]int{}
	m := make(map[string]bool)

	l := len(nums)
	sort.Ints(nums)
	for i1 := 0; i1 < l-3; i1++ {
		for i2 := i1 + 1; i2 < l-2; i2++ {
			for i3 := i2 + 1; i3 < l-1; i3++ {
				n := target - nums[i1] - nums[i2] - nums[i3]
				for i4 := l - 1; i4 > i3 && nums[i4] >= n; i4-- {
					if nums[i4] == n {
						d := fmt.Sprint(nums[i1], nums[i2], nums[i3], nums[i4])
						if _, ok := m[d]; !ok {
							m[d] = true
							ans = append(ans, []int{nums[i1], nums[i2], nums[i3], nums[i4]})
						}
					}
				}
			}
		}
	}

	return ans

}
