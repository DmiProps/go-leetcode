package problems

import "fmt"

// Problem20 - 20. Valid Parentheses
func Problem20() {

	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

}

func isValid(s string) bool {

	a := []byte{}

	for _, c := range s {
		l := len(a)
		switch {
		case c == '(' || c == '{' || c == '[':
			a = append(a, byte(c))
		case l == 0:
			return false
		case c == ')' && a[l-1] == '(':
			a = a[:l-1]
		case c == ']' && a[l-1] == '[':
			a = a[:l-1]
		case c == '}' && a[l-1] == '{':
			a = a[:l-1]
		default:
			return false
		}
	}

	return len(a) == 0

}
