package problems

import "fmt"

// ProblemDay4 - run test problem
func ProblemDay4() {

	test1 := []int{0, 1, 0, 3, 12}

	moveZeroes(test1)

	fmt.Println(test1)

	test2 := []int{0, 1, 0}

	moveZeroes(test2)

	fmt.Println(test2)

}

func moveZeroes(nums []int) {

	cnt := 0

	for idx := 0; idx < len(nums); idx++ {

		if nums[idx] == 0 {
			cnt++
		} else if cnt > 0 {
			nums[idx-cnt] = nums[idx]
			nums[idx] = 0
		}

	}

}
