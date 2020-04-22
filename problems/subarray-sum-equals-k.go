package problems

import "fmt"

// Problem560 - 560. Subarray Sum Equals K
func Problem560() {

	fmt.Println(subarraySum([]int{1, 1, 1}, 2))

}

func subarraySum(nums []int, k int) int {

	s := make([]int, len(nums)+1)
	a := 0
	for idx, n := range nums {
		a += n
		s[idx+1] = a
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i]-s[j] == k {
				ans++
			}
		}
	}
	return ans

}
