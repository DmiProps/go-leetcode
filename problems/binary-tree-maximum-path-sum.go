package problems

import "fmt"

// Problem124 - 124. Binary Tree Maximum Path Sum
func Problem124() {

	/*test1 := &TreeNode1{Val: 1}
	test1.Left = &TreeNode1{Val: 2}
	test1.Right = &TreeNode1{Val: 3}

	fmt.Println(maxPathSum(test1))

	test2 := &TreeNode1{Val: -10}
	test2.Left = &TreeNode1{Val: 9}
	test2.Right = &TreeNode1{Val: 20}
	test2.Right.Left = &TreeNode1{Val: 15}
	test2.Right.Right = &TreeNode1{Val: 7}

	fmt.Println(maxPathSum(test2))

	test3 := &TreeNode1{Val: -3}

	fmt.Println(maxPathSum(test3))

	test4 := &TreeNode1{Val: -2}
	test4.Left = &TreeNode1{Val: -1}

	fmt.Println(maxPathSum(test4))

	test5 := &TreeNode1{Val: -8}
	test5.Left = &TreeNode1{Val: -9}
	test5.Right = &TreeNode1{Val: -2}
	test5.Right.Left = &TreeNode1{Val: 3}
	test5.Right.Left.Right = &TreeNode1{Val: 8}

	fmt.Println(maxPathSum(test5))

	test6 := &TreeNode1{Val: 5}
	test6.Left = &TreeNode1{Val: 4}
	test6.Right = &TreeNode1{Val: 8}

	test6.Left.Left = &TreeNode1{Val: 11}
	test6.Left.Left.Left = &TreeNode1{Val: 7}
	test6.Left.Left.Right = &TreeNode1{Val: 2}

	test6.Right.Left = &TreeNode1{Val: 13}
	test6.Right.Right = &TreeNode1{Val: 4}
	test6.Right.Right.Right = &TreeNode1{Val: 1}

	fmt.Println(maxPathSum(test6))

	test7 := &TreeNode1{Val: -3}
	test7.Left = &TreeNode1{Val: 2}
	test7.Left.Left = &TreeNode1{Val: 3}
	test7.Left.Right = &TreeNode1{Val: -6}
	test7.Left.Right.Left = &TreeNode1{Val: 2}

	fmt.Println(maxPathSum(test7))*/

	test8 := &TreeNode1{Val: -10}
	test8.Left = &TreeNode1{Val: 9}
	test8.Right = &TreeNode1{Val: 20}

	test8.Right.Left = &TreeNode1{Val: 15}
	test8.Right.Right = &TreeNode1{Val: 7}

	fmt.Println(maxPathSum(test8))

}

// TreeNode1 - struct
type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode1) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	l1, l2 := maxPathSumTest(root.Left)
	r1, r2 := maxPathSumTest(root.Right)

	ans := root.Val
	if root.Left != nil && l1 > 0 {
		ans += l1
	}
	if root.Right != nil && r1 > 0 {
		ans += r1
	}

	if root.Left != nil && l2 > ans {
		ans = l2
	}
	if root.Right != nil && r2 > ans {
		ans = r2
	}

	return ans

}

func maxPathSumTest(n *TreeNode1) (int, int) {

	// ans1 - with node
	// ans2 - without node

	if n == nil {
		return 0, 0
	}
	l1, l2 := maxPathSumTest(n.Left)
	r1, r2 := maxPathSumTest(n.Right)

	a := 0

	if n.Left != nil && l1 > 0 {
		a = l1
	}
	if n.Right != nil && r1 > 0 {
		a = max(a, r1)
	}
	ansWithNode := n.Val + a

	ansWithoutNode := ansWithNode

	if n.Left != nil && l2 > ansWithoutNode {
		ansWithoutNode = l2
	}
	if n.Right != nil && r2 > ansWithoutNode {
		ansWithoutNode = r2
	}
	if n.Left != nil && n.Right != nil {
		ansWithoutNode = max(ansWithoutNode, l1+r1+n.Val)
	}

	return ansWithNode, ansWithoutNode

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
