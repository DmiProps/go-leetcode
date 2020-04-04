package problems

import (
	"fmt"
	"sort"
)

// Biweekly23_4 - test 4
func Biweekly23_4() {

	test1 := []int{-1, -8, 0, 5, -9}

	fmt.Println(maxSatisfaction(test1))

	test2 := []int{4, 3, 2}

	fmt.Println(maxSatisfaction(test2))

	test3 := []int{-1, -4, -5}

	fmt.Println(maxSatisfaction(test3))

	test4 := []int{-2, 5, -1, 0, 3, -3}

	fmt.Println(maxSatisfaction(test4))

}

func maxSatisfaction(satisfaction []int) int {

	sort.Ints(satisfaction)

	sum := 0
	cnt := 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		sum += satisfaction[i]
		if sum < 0 {
			cnt = len(satisfaction) - i - 1
			break
		} else {
			cnt = len(satisfaction) - i
		}

	}
	if cnt == 0 {
		return 0
	}
	sum = 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		sum += cnt * satisfaction[i]
		cnt--
		if cnt == 0 {
			break
		}
	}
	return sum

}
