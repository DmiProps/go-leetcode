package problems

import "fmt"

// Problem238 - 238. Product of Array Except Self
func Problem238() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{1, 0}))

}

func productExceptSelf(nums []int) []int {

	z := 0
	x := 1
	for _, n := range nums {
		if n == 0 {
			z++
		} else {
			x *= n
		}
	}
	ans := make([]int, len(nums))
	if z > 1 {
		return ans
	}
	for idx, n := range nums {
		if z > 0 {
			if n == 0 {
				ans[idx] = x
			} else {
				ans[idx] = 0
			}
		} else {
			ans[idx] = x / n
		}
	}

	return ans

}
