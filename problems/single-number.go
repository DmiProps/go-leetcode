package problems

import "fmt"

// ProblemDay1 - 136. Single Number
func ProblemDay1() {

	test1 := []int{2, 2, 1}
	fmt.Println(singleNumber(test1))

	test2 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(test2))

}

func singleNumber(nums []int) int {

	n := make(map[int]bool, len(nums))

	for _, i := range nums {
		_, ok := n[i]
		if !ok {
			n[i] = false
		} else {
			n[i] = true
		}
	}
	for i, ok := range n {
		if !ok {
			return i
		}
	}
	return 0
}
