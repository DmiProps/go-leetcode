package problems

import (
	"fmt"
	"sort"
)

// Problem16 - 16. 3Sum Closest
func Problem16() {

	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{13, 2, 0, -14, -20, 19, 8, -5, -13, -3, 20, 15, 20, 5, 13, 14, -17, -7, 12, -6, 0, 20, -19, -1, -15, -2, 8, -2, -9, 13, 0, -3, -18, -9, -9, -19, 17, -14, -19, -4, -16, 2, 0, 9, 5, -7, -4, 20, 18, 9, 0, 12, -1, 10, -17, -11, 16, -13, -14, -3, 0, 2, -18, 2, 8, 20, -15, 3, -13, -12, -2, -19, 11, 11, -10, 1, 1, -10, -2, 12, 0, 17, -19, -7, 8, -19, -17, 5, -5, -10, 8, 0, -12, 4, 19, 2, 0, 12, 14, -9, 15, 7, 0, -16, -5, 16, -12, 0, 2, -16, 14, 18, 12, 13, 5, 0, 5, 6}, -59))

}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	l := len(nums)
	d := 0
	ans := 0
	f := false
	for idx1 := 0; idx1 < l-2; idx1++ {

		for idx2 := idx1 + 1; idx2 < l-1; idx2++ {
			n := target - nums[idx1] - nums[idx2]
			lastDist := 0
			for idx3 := l - 1; idx3 > idx2; idx3-- {
				d1 := dist(nums[idx3], n)
				if idx3 == l-1 {
					lastDist = d1
					if !f {
						f = true
						d = lastDist
						ans = nums[idx1] + nums[idx2] + nums[idx3]
					}
				} else {
					if d1 > lastDist {
						break
					}
				}
				if d1 < d {
					d = d1
					ans = nums[idx1] + nums[idx2] + nums[idx3]
				}
				lastDist = d1

			}
		}

	}
	return ans

}

func dist(n, m int) int {
	d := m - n
	if d < 0 {
		return -d
	}
	return d
}
