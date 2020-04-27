package problems

import "fmt"

// Problem221 - 221. Maximal Square
func Problem221() {

	test1 := [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0}}

	fmt.Println(maximalSquare(test1))

}

func maximalSquare(matrix [][]byte) int {

	ans := 0
	ly := len(matrix)
	if ly == 0 {
		return 0
	}
	lx := len(matrix[0])
	if lx == 0 {
		return 0
	}

	for y := 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if matrix[y][x] == 1 {
				max := 1
				for s, f := 1, true; s+y < ly && s+x < lx && f; s++ {
					for d := 0; d < s; d++ {
						if matrix[y+d][x+s] != 1 || matrix[y+s][x+d] != 1 {
							f = false
							break
						}
					}
					if f && matrix[y+s][x+s] != 1 {
						f = false
					}
					if f {
						max = s + 1
					}
				}
				if max*max > ans {
					ans = max * max
				}
			}
		}
	}

	return ans

}
