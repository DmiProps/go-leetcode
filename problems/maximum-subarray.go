package problems

import "fmt"

// ProblemDay3 - 53. Maximum Subarray
func ProblemDay3() {

	test1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(test1))

	test2 := []int{-2, -3, -1, -5, -4}

	fmt.Println(maxSubArray(test2))

	test3 := []int{-1, 0, -2}

	fmt.Println(maxSubArray(test3))

}

func maxSubArray(nums []int) int {

	ans := 0
	sum := 0
	max := 1

	for _, n := range nums {
		sum += n
		if max == 1 || max < n {
			max = n
		}
		if sum < 0 {
			sum = 0
		} else if sum > ans {
			ans = sum
		}
	}
	if ans == 0 {
		ans = max
	}
	return ans

}
