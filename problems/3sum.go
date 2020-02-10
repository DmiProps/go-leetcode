package problems

import (
	"fmt"
	"sort"
)

// Problem15 run test problem
func Problem15() {

	test := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(test)

	fmt.Println(res)

}

func threeSum(nums []int) [][]int {

	res := make([][]int, 0, 10)
	sort.Ints(nums)
	s := 0
	l := 0
	a := false
	for i1 := 0; i1 < len(nums)-2; i1++ {
		if nums[i1] > 0 {
			break
		}
		for i2 := i1 + 1; i2 < len(nums)-1; i2++ {
			if nums[i1]+nums[i2] > 0 {
				break
			}
			for i3 := i2 + 1; i3 < len(nums); i3++ {
				s = nums[i1] + nums[i2] + nums[i3]
				if s == 0 {
					a = true
					if l > 0 {
						if res[l-1][0] == nums[i1] && res[l-1][1] == nums[i2] && res[l-1][2] == nums[i3] {
							a = false
						}
					}
					if a {
						res = append(res, []int{nums[i1], nums[i2], nums[i3]})
						l++
					}
				} else if s > 0 {
					break
				}
			}
		}
	}
	return res

}
