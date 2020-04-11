package problems

import "fmt"

// Problem48 - 48. Rotate Image
func Problem48() {

	test1 := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	rotateMatrix(test1)
	fmt.Println(test1)

	test2 := [][]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}

	rotateMatrix(test2)
	fmt.Println(test2)

}

func rotateMatrix(matrix [][]int) {

	l := len(matrix)
	a := 0
	for contur := 0; contur < l/2; contur++ {
		for pos := 0; pos < l-1-contur*2; pos++ {
			a = matrix[contur][contur+pos]
			matrix[contur][contur+pos] = matrix[l-1-contur-pos][contur]
			matrix[l-1-contur-pos][contur] = matrix[l-1-contur][l-1-contur-pos]
			matrix[l-1-contur][l-1-contur-pos] = matrix[contur+pos][l-1-contur]
			matrix[contur+pos][l-1-contur] = a
		}
	}

}
