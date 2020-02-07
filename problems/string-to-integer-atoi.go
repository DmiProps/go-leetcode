package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// Problem8 run test problem
func Problem8() {

	tests := [...]string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
	}
	for _, test := range tests {
		res := myAtoi(test)
		fmt.Printf("In: %s, out: %d\n", test, res)
	}

}

func myAtoi(str string) int {

	str = strings.TrimLeft(str, " ")

	var digits = "0123456789"
	var end = len(str)

	for idx, ch := range str {
		if idx == 0 && (ch == '-' || ch == '+') {
			continue
		}
		if strings.ContainsRune(digits, ch) {
			continue
		}
		end = idx
		break
	}
	str = str[:end]

	i, _ := strconv.Atoi(str)
	if i < -2147483648 {
		i = -2147483648
	} else if i > 2147483647 {
		i = 2147483647
	}
	return i

}
