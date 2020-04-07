package problems

import "fmt"

// ProblemPermute - phone mck problem
func ProblemPermute() {

	//test1 := []int{1, 2, 3}
	//fmt.Println(permute(test1))

	test2 := []int{5, 4, 6, 2}
	fmt.Println(permute(test2))

}

func permute(nums []int) [][]int {
	cnt := 1
	l := len(nums)
	if l == 0 {
		return [][]int{{}}
	}
	if l == 1 {
		return [][]int{{nums[0]}}
	}
	for i := 2; i <= l; i++ {
		cnt *= i
	}
	ans := make([][]int, cnt)
	fillNums(ans, 0, nums, cnt, l, 0)
	return ans
}

func fillNums(ans [][]int, start int, nums []int, cnt int, l int, def int) {
	for idx := 0; idx < l; idx++ {
		newNums := make([]int, l-1)
		if l > 2 {
			for i := 0; i < l; i++ {
				if i == idx {
					continue
				}
				if i < idx {
					newNums[i] = nums[i]
				} else {
					newNums[i-1] = nums[i]
				}
			}
		}
		for i := 0; i < cnt/l; i++ {
			if ans[start+i+idx*cnt/l] == nil {
				ans[start+i+idx*cnt/l] = make([]int, l)
			}
			ans[start+i+idx*cnt/l][def] = nums[idx]
			if l == 2 {
				ans[start+i+idx*cnt/l][def+1] = nums[1-idx]
			}
		}
		if l > 2 {
			fillNums(ans, start+idx*cnt/l, newNums, cnt/l, l-1, def+1)
		}
	}
}
