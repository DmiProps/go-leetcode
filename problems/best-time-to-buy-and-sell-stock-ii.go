package problems

import "fmt"

// ProblemDay5 - 122. Best Time to Buy and Sell Stock II
func ProblemDay5() {

	test1 := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(test1))

	test2 := []int{1, 2, 3, 4, 5}

	fmt.Println(maxProfit(test2))

	test3 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(test3))

}

func maxProfit(prices []int) int {

	ans := 0
	dif := 0

	for idx := 0; idx < len(prices)-1; idx++ {
		dif = prices[idx+1] - prices[idx]
		if dif > 0 {
			ans += dif
		}
	}

	return ans

}
