package problems

import (
	"fmt"
	"sort"
)

// Problem15 - 15. 3Sum
func Problem15() {

	//test := []int{-1, 0, 1, 2, -1, -4}
	//test := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	test := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	res := threeSum(test)

	fmt.Println(res)

}

func testArray(i [][]int, j1, j2, j3 int) bool {

	for a := len(i) - 1; a >= 0; a-- {
		if i[a][0] == j1 && i[a][1] == j2 && i[a][2] == j3 {
			return false
		}
	}
	return true

}

func threeSum(nums []int) [][]int {

	n := make([]int, 0, len(nums))
	sort.Ints(nums)
	j := 0
	for i := 0; i < len(nums); i++ {
		if j > 2 {
			if n[j-1] == nums[i] && n[j-2] == nums[i] && n[j-3] == nums[i] {
				continue
			}
		}
		n = append(n, nums[i])
		j++
	}

	res := make([][]int, 0, 10)

	s := 0
	l := 0
	a := false
	for i1 := 0; i1 < len(n)-2; i1++ {
		if n[i1] > 0 {
			break
		}
		for i2 := i1 + 1; i2 < len(n)-1; i2++ {
			if n[i1]+n[i2] > 0 {
				break
			}
			for i3 := i2 + 1; i3 < len(n); i3++ {
				s = n[i1] + n[i2] + n[i3]
				if s == 0 {
					a = true
					if l > 0 {
						if !testArray(res, n[i1], n[i2], n[i3]) {
							a = false
						}
					}
					if a {
						res = append(res, []int{n[i1], n[i2], n[i3]})
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
