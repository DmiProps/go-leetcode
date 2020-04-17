package problems

import "fmt"

// Problem200 - 200. Number of Islands
func Problem200() {

	test1 := [][]byte{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0}} // 1
	fmt.Println(numIslands(test1))

	test2 := [][]byte{
		{1,1,0,0,0},
		{1,1,0,0,0},
		{0,0,1,0,0},
		{0,0,0,1,1}} // 3
	fmt.Println(numIslands(test2))

	test3 := [][]byte{
		{1,0,1,0,0},
		{1,1,1,0,0},
		{0,0,1,0,1},
		{0,0,0,1,1}} // 2
	fmt.Println(numIslands(test3))

}

func numIslands(grid [][]byte) int {

	ans := 0

	for y, h := range grid {
		for x, n := range h {
			if n == 1 {
				ans++
				fillIsland(grid, x, y)
			}
		}
	}

	return ans
    
}

func fillIsland(grid [][]byte, x int, y int) {

	grid[y][x] = 2
	if y > 0 && grid[y-1][x] == 1 {
		fillIsland(grid, x, y-1)
	}
	if x > 0 && grid[y][x-1] == 1 {
		fillIsland(grid, x-1, y)
	}
	if y < len(grid) - 1 && grid[y+1][x] == 1 {
		fillIsland(grid, x, y+1)
	}
	if x < len(grid[y]) - 1 && grid[y][x+1] == 1 {
		fillIsland(grid, x+1, y)
	}

}