package problems

import (
	"fmt"
	"sort"
)

// Problem350 - 350. Intersection of Two Arrays II
func Problem350() {

	test11 := []int{1, 2, 2, 1}
	test12 := []int{2, 2}
	fmt.Println(intersect(test11, test12))

	test21 := []int{4, 9, 5}
	test22 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(test21, test22))

}

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)
	l1 := len(nums1)
	l2 := len(nums2)
	ans := []int{}
	for idx1, idx2 := 0, 0; idx1 < l1 && idx2 < l2; {
		if nums1[idx1] == nums2[idx2] {
			ans = append(ans, nums1[idx1])
			idx1++
			idx2++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			idx1++
		}
	}
	return ans

}
