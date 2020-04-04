package problems

import "fmt"

// Biweekly23_1 - test 1
func Biweekly23_1() {

	fmt.Println(countLargestGroup(13))
	fmt.Println(countLargestGroup(2))
	fmt.Println(countLargestGroup(15))
	fmt.Println(countLargestGroup(24))

}

func countLargestGroup(n int) int {

	max := 0
	grp := make(map[int]int)

	for i := 1; i <= n; i++ {
		s := 0
		p := i
		for p > 0 {
			s += p % 10
			p /= 10
		}
		cnt := grp[s] + 1
		if cnt > max {
			max = cnt
		}
		grp[s] = cnt
	}
	cnt := 0
	for _, val := range grp {
		if val == max {
			cnt++
		}
	}
	return cnt

}
