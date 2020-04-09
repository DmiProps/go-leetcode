package problems

import (
	"fmt"
)

// Problem844 - 844. Backspace String Compare
func Problem844() {

	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a#c", "b"))

}

func backspaceCompare(S string, T string) bool {

	idx1 := len(S) - 1
	idx2 := len(T) - 1
	del := 0
	for {
		del = 0
		for idx1 >= 0 {
			if S[idx1] == '#' {
				del++
			} else {
				del--
			}
			if del < 0 {
				break
			}
			idx1--
		}
		del = 0
		for idx2 >= 0 {
			if T[idx2] == '#' {
				del++
			} else {
				del--
			}
			if del < 0 {
				break
			}
			idx2--
		}
		if idx1 < 0 && idx2 < 0 {
			return true
		}
		if idx1 < 0 || idx2 < 0 {
			return false
		}
		if S[idx1] != T[idx2] {
			return false
		}
		idx1--
		idx2--
	}

}
