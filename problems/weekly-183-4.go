package problems

import "fmt"

// Weekly183_4 - 1406. Stone Game III
func Weekly183_4() {

	fmt.Println("Task 4...")

	test1 := []int{1, 2, 3, 7}
	fmt.Println(stoneGameIII(test1))

	test2 := []int{1, 2, 3, -9}
	fmt.Println(stoneGameIII(test2))

	test3 := []int{1, 2, 3, 6}
	fmt.Println(stoneGameIII(test3))

	test4 := []int{1, 2, 3, -1, -2, -3, 7}
	fmt.Println(stoneGameIII(test4))

	test5 := []int{-1, -2, -3}
	fmt.Println(stoneGameIII(test5))

}

func stoneGameIII(stoneValue []int) string {

	ans := 0
	//max := 0

	for idx := len(stoneValue) - 1; idx >= 0; idx-- {

		if stoneValue[idx] < 0 {
			ans = -1
		} else if stoneValue[idx] > 0 {
			ans = 1
		} else {
			ans = 0
		}
		//max = stoneValue[idx]

	}
	if ans < 0 {
		return "Bob"
	}
	if ans > 0 {
		return "Alice"
	}
	return "Tie"

}
