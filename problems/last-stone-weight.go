package problems

import (
	"fmt"
	"sort"
)

// Problem1046 - 1046. Last Stone Weight
func Problem1046() {

	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

}

func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		l := len(stones)
		stones = append(stones[:l-2], stones[l-1]-stones[l-2])
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]

}
