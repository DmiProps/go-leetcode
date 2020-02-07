package problems

import "fmt"

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

	return 0

}
