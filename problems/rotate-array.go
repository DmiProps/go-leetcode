package problems

import "fmt"

// Problem189 - 189. Rotate Array
func Problem189() {

	test1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(test1, 3)
	fmt.Println(test1)

	test2 := []int{-1, -100, 3, 99}
	rotate(test2, 2)
	fmt.Println(test2)

}

func rotate(nums []int, k int) {

	l := len(nums)
	if l == 0 || k%l == 0 {
		return
	}
	k %= l

	buf := make([]int, k)
	for i := 0; i < k; i++ {
		buf[i] = nums[l-k+i]
	}
	for i := l - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = buf[i]
	}

}
