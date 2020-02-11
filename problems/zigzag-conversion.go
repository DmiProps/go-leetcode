package problems

import "fmt"

// Problem6 is runs solution of problem and shows result.
func Problem6() {

	fmt.Println("Got:\n", convert("PAYPALISHIRING", 3))
	fmt.Println("Expected:\n", "PAHNAPLSIIGYIR")

	fmt.Println("Got:\n", convert("PAYPALISHIRING", 4))
	fmt.Println("Expected:\n", "PINALSIGYAHRPI")

	fmt.Println("Got:\n", convert("A", 1))
	fmt.Println("Expected:\n", "A")

	fmt.Println("Got:\n", convert("AB", 1))
	fmt.Println("Expected:\n", "AB")

}

func convert(s string, numRows int) string {

	if s == "" {
		return ""
	}

	l := len(s)
	ans := make([]byte, l, l)
	numCols := l
	if numRows > 1 {
		numCols = (l-1)/(2*numRows-2) + 1
	}
	i := 0

	step := 1
	if numRows > 1 {
		step = 2 * (numRows - 1)
	}

	for row := 0; row < numRows; row++ {
		switch {
		case row == 0:
			for col := 0; col < numCols; col++ {
				ans[i] = s[col*step]
				i++
			}
		case row == numRows-1:
			for col := 0; col < numCols; col++ {
				if col*step+numRows-1 >= l {
					break
				}
				ans[i] = s[col*step+numRows-1]
				i++
			}
		default:
			for col := 0; col < numCols; col++ {
				if col*step+row >= l {
					break
				}
				ans[i] = s[col*step+row]
				i++
				if col*step+step-row >= l {
					break
				}
				ans[i] = s[col*step+step-row]
				i++
			}
		}
	}

	return string(ans)

}
