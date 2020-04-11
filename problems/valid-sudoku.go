package problems

import "fmt"

// Problem36 - 36. Valid Sudoku
func Problem36() {

	test1 := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(test1))

	test2 := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(test2))

}

func isValidSudoku(board [][]byte) bool {

	s := []map[byte]bool{nil, nil, nil, nil, nil, nil, nil, nil, nil}
	for idx := 0; idx < 9; idx++ {
		s[idx] = make(map[byte]bool, 9)
	}

	idx := 0
	for idx1 := 0; idx1 < 9; idx1++ {
		h := make(map[byte]bool, 9)
		v := make(map[byte]bool, 9)
		for idx2 := 0; idx2 < 9; idx2++ {
			if board[idx1][idx2] != '.' {
				if _, ok := h[board[idx1][idx2]]; ok {
					return false
				}
				h[board[idx1][idx2]] = true
				idx = (idx1/3)*3 + idx2/3
				if _, ok := s[idx][board[idx1][idx2]]; ok {
					return false
				}
				s[idx][board[idx1][idx2]] = true
			}
			if board[idx2][idx1] != '.' {
				if _, ok := v[board[idx2][idx1]]; ok {
					return false
				}
				v[board[idx2][idx1]] = true
			}
		}
	}
	return true

}
