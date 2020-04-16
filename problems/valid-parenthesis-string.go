package problems

import "fmt"

// Problem678 - 678. Valid Parenthesis String
func Problem678() {

	fmt.Println(checkValidString("()"))                                                // true
	fmt.Println(checkValidString("(*)"))                                               // true
	fmt.Println(checkValidString("(*))"))                                              // true
	fmt.Println(checkValidString("()())"))                                             // false
	fmt.Println(checkValidString("()()"))                                              // true
	fmt.Println(checkValidString("()())*"))                                            // false
	fmt.Println(checkValidString("()()*"))                                             // true
	fmt.Println(checkValidString("*())"))                                              // true
	fmt.Println(checkValidString("((**"))                                              // true
	fmt.Println(checkValidString("(()"))                                               // false
	fmt.Println(checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*")) // false

}

func checkValidString(s string) bool {

	cnt := 0
	lev := 0
	for _, ch := range s {
		switch {
		case ch == '(':
			cnt++
		case ch == '*':
			lev++
		case cnt > 0:
			cnt--
		case lev > 0:
			lev--
		default:
			return false
		}
	}
	cnt = 0
	lev = 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		switch {
		case ch == ')':
			cnt++
		case ch == '*':
			lev++
		case cnt > 0:
			cnt--
		case lev > 0:
			lev--
		default:
			return false
		}
	}

	return true

}
