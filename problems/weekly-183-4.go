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

	a := make([]int, 3)
	b := make([]int, 3)

	isTie := false

	for i := 0; i < 3 && i < len(stoneValue); i++ {
		a[i], b[i] = stoneGameIIIBob(stoneValue, i+1)
		if a[i] > b[i] {
			return "Alice"
		} else if a[i] == b[i] {
			isTie = true
		}
	}
	if isTie {
		return "Tie"
	}
	return "Bob"

}

func stoneGameIIIBob(stoneValue []int, pos int) (int, int) {

	/*a := make([]int, 3)
	b := make([]int, 3)

	isTie := false

	for i := pos; i < 3 + pos && i < len(stoneValue); i++ {
		a[i], b[i] = stoneGameIIIAlice(stoneValue, i+1)
		if b[i] > a[i] {
			return "Alice"
		} else if a[i] == b[i] {
			isTie = true
		}
	}
	if isTie {
		return "Tie"
	}
	return "Bob"*/
	return 0, 0

}
