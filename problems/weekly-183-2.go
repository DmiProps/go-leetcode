package problems

import "fmt"

// Weekly183_2 - 1404. Number of Steps to Reduce a Number in Binary Representation to One
func Weekly183_2() {

	fmt.Println("Task 2...")

	fmt.Println(numSteps("1101"))
	fmt.Println(numSteps("10"))
	fmt.Println(numSteps("1"))
	fmt.Println(numSteps("1001"))
	fmt.Println(numSteps("1111110011101010110011100100101110010100101110111010111110110010"))

}

func numSteps(s string) int {

	ans := 0
	bits := make([]bool, len(s)+1)

	for idx, ch := range s {
		bits[idx+1] = ch == rune('1')
	}
	pos := len(s)
	for pos > 0 {
		if pos == 1 && !bits[pos-1] {
			break
		}
		if bits[pos] {
			for idx := pos; ; idx-- {
				if !bits[idx] {
					bits[idx] = true
					break
				} else {
					bits[idx] = false
				}
			}
		} else {
			pos--
		}
		ans++
	}

	return ans

}
