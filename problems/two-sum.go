package problems

import "fmt"

// Problem1 - 1. Two Sum
func Problem1() {

	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}

func twoSum(nums []int, target int) []int {

	m := make(map[int][]int, len(nums))
	for idx, n := range nums {
		if v, ok := m[n]; ok {
			v[1] = idx
		} else {
			m[n] = []int{idx, -1}
		}
	}

	for k, v := range m {
		if k*2 == target {
			if v[1] >= 0 {
				return v
			}
			continue
		}
		if v1, ok := m[target-k]; ok {
			return []int{v[0], v1[0]}
		}
	}

	return []int{0, 0}

}
