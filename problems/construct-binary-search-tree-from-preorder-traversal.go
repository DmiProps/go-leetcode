package problems

import "fmt"

// Problem1008 - 1008. Construct Binary Search Tree from Preorder Traversal
func Problem1008() {

	/*test1 := TreeNode2{Val: 8}
	test1.Left = TreeNode2{Val: 5}
	test1.Left.Left = TreeNode2{Val: 1}
	test1.Left.Right = TreeNode2{Val: 7}
	test1.Right = TreeNode2{Val: 10}
	test1.Right.Right = TreeNode2{Val: 12}*/

	ans := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})

	fmt.Println(ans)

}

// TreeNode2 - tree
type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode2 {

	l := len(preorder)
	if l == 0 {
		return nil
	}
	ans := &TreeNode2{Val: preorder[0]}

	idx, left := fromPreorder(preorder, ans, 1)
	ans.Left = left

	if idx < l {
		_, right := fromPreorder(preorder, nil, idx)
		ans.Right = right
	}

	return ans

}

func fromPreorder(p []int, top *TreeNode2, start int) (int, *TreeNode2) {

	if start >= len(p) {
		return start, nil
	}
	if top != nil && p[start] > top.Val {
		return start, nil
	}

	ans := &TreeNode2{Val: p[start]}

	idx, left := fromPreorder(p, ans, start+1)
	ans.Left = left

	if idx < len(p) {
		i, right := fromPreorder(p, top, idx)
		idx = i
		ans.Right = right
	}
	return idx, ans

}
