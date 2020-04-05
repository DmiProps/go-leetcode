package problems

import "fmt"

// Weekly183_3 - 1405. Longest Happy String
func Weekly183_3() {

	fmt.Println("Task 3...")

	fmt.Println(longestDiverseString(1, 1, 7))
	fmt.Println(longestDiverseString(2, 2, 1))
	fmt.Println(longestDiverseString(7, 1, 0))
	fmt.Println(longestDiverseString(0, 8, 11))

}

func longestDiverseString(a int, b int, c int) string {

	ans := ""
	last := ""

	for a+b+c > 0 {
		if a > 0 && last != "a" && (a >= b || last == "b") && (a >= c || last == "c") {
			ans += "a"
			a--
			if a > 0 && !(b > a || c > a) {
				ans += "a"
				a--
			}
			last = "a"
		} else if b > 0 && last != "b" && (b >= a || last == "a") && (b >= c || last == "c") {
			ans += "b"
			b--
			if b > 0 && !(a > b || c > b) {
				ans += "b"
				b--
			}
			last = "b"
		} else if c > 0 && last != "c" && (c >= a || last == "a") && (c >= b || last == "b") {
			ans += "c"
			c--
			if c > 0 && !(a > c || b > c) {
				ans += "c"
				c--
			}
			last = "c"
		} else {
			break
		}
	}

	return ans

}
