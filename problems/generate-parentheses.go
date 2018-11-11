package problems

import "fmt"

// Problem22 is runs solution of problem and shows result.
func Problem22() {

	/* variant 1:
	[
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"
	]*/
	n := 3

	fmt.Println(generateParenthesis(n))

}

// Problem #22 (medium level): Generate Parentheses.
func generateParenthesis(n int) []string {
	
	ans := make([]string, 0)
	if n == 0 {
		return ans
	}
	if n == 1 {
		ans = append(ans, "()")
		return ans
	}

	var min, max int
	for i := 0; i < 2 * n; i++ {
		
		max <<= 1
		if i < n {
			max++
		}
		
		min += i % 2
		min <<= 1
		
	}

	for i := min; i <= max; i += 2 {
		if s, ok := checkParentheses(i, 2 * n); ok {
			ans = append(ans, s)
		}
	}

	return ans
}

func checkParentheses(i, size int) (s string, ok bool) {

	b := make([]byte, size, size)
	c := 0

	for pos := 0; i > 0; i >>= 1 {
		c += 1 - 2 * (i % 2)
		if c < 0 {
			return
		}
		if i % 2 == 0 {
			b[size - pos - 1] = ')'
		} else {
			b[size - pos - 1] = '('
		}
		pos++
	}
	if c == 0 {
		ok = true
		s = string(b)
	}
	return
}