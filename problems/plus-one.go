package problems

import "fmt"

// Problem66 - 66. Plus One
func ProblemX() {
	test1 := []int{1, 2, 3}
	fmt.Println(plusOne(test1))
	test2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(test2))
	test3 := []int{9, 9}
	fmt.Println(plusOne(test3))
	test4 := []int{0}
	fmt.Println(plusOne(test4))
	test5 := []int{8, 9, 9, 9}
	fmt.Println(plusOne(test5))
}

func plusOne(digits []int) []int {
	allNine := true
	for _, n := range digits {
		if n != 9 {
			allNine = false
			break
		}
	}
	if allNine {
		ans := make([]int, len(digits)+1)
		ans[0] = 1
		return ans
	}
	p := 1
	for idx := len(digits) - 1; idx >= 0; idx-- {
		d := digits[idx] + p
		digits[idx] = d % 10
		p = d / 10
		if p == 0 {
			break
		}
	}
	return digits
}
