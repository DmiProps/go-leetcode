package problems

import "fmt"

// ProblemX1 - Counting Elements
func ProblemX1() {

	fmt.Println(countElements([]int{1, 1, 2}))
	fmt.Println(countElements([]int{1, 2, 3}))
	fmt.Println(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))
	fmt.Println(countElements([]int{1, 3, 2, 3, 5, 0}))
	fmt.Println(countElements([]int{1, 1, 2, 2}))

}

func countElements(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	m := make(map[int]int, 1)
	for _, n := range arr {
		m[n]++
	}
	ans := 0
	for v, c := range m {
		if m[v+1] > 0 {
			ans += c
		}
	}
	return ans
}
