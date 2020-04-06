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

	test6 := []int{11, 10, 7, -6, -17, -5, 6, -1, 6, -13, 13, 2, 4, -16, 13, 13, 4, 0, 16, 16, -16, 5, -16}
	fmt.Println(stoneGameIII(test6))

}

func stoneGameIII(stoneValue []int) string {

	l := len(stoneValue)

	sum := make([]int, l+1)
	for idx := l - 1; idx >= 0; idx-- {
		sum[idx] = sum[idx+1] + stoneValue[idx]
	}

	max := make([]int, l)
	for idx := l - 1; idx >= 0 && idx > l-4; idx-- {
		max[idx] = sum[idx]
	}

	for idx := l - 4; idx >= 0; idx-- {
		min := 0
		for j := idx + 1; j < l && j < idx+4; j++ {
			if j == idx+1 || max[j] < min {
				min = max[j]
			}
		}
		max[idx] = sum[idx] - min
	}

	a := max[0]
	//b :=

	for i := 1; i < 3 && i < l; i++ {
		if max[i] > a {
			a = max[i]
		}
	}

	b := sum[0] - a
	if a == b {
		return "Tie"
	}
	if a > b {
		return "Alice"
	}
	return "Bob"

}
