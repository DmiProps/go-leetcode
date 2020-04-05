package problems

import "fmt"

// ProblemDay2 - 202. Happy Number
func ProblemDay2() {

	fmt.Println(2, isHappy(2))

	for i := 1; i <= 10000; i++ {
		fmt.Println(i, isHappy(i))
	}

}

func isHappy(n int) bool {

	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n == 2 || n == 4 {
			return false
		}
		r := 0
		for n != 0 {
			r += (n % 10) * (n % 10)
			n /= 10
		}
		n = r
	}

}
