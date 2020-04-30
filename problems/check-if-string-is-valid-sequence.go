package problems

import "fmt"

// ProblemX3 - Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree
func ProblemX3() {

	test1 := &TreeNode3{Val: 0}
	test1.Left = &TreeNode3{Val: 1}
	test1.Right = &TreeNode3{Val: 0}
	test1.Left.Left = &TreeNode3{Val: 0}
	test1.Left.Right = &TreeNode3{Val: 1}
	test1.Right.Left = &TreeNode3{Val: 0}
	test1.Left.Left.Right = &TreeNode3{Val: 1}
	test1.Left.Right.Left = &TreeNode3{Val: 0}
	test1.Left.Right.Left = &TreeNode3{Val: 0}

	fmt.Println(isValidSequence(test1, []int{0, 1, 0, 1}))

	test2 := &TreeNode3{Val: 8}
	test2.Left = &TreeNode3{Val: 3}
	test2.Left.Left = &TreeNode3{Val: 2}
	test2.Left.Right = &TreeNode3{Val: 1}
	test2.Left.Left.Left = &TreeNode3{Val: 5}
	test2.Left.Left.Right = &TreeNode3{Val: 4}

	fmt.Println(isValidSequence(test2, []int{8}))
}

// TreeNode3 - Definition for a binary tree node.
type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidSequence(root *TreeNode3, arr []int) bool {

	if root == nil {
		return false
	}

	isLeaf := (root.Left == nil && root.Right == nil)
	if isLeaf {
		if len(arr) != 1 {
			return false
		}
		if root.Val != arr[0] {
			return false
		}
		return true
	}
	if len(arr) < 2 {
		return false
	}
	if arr[0] != root.Val {
		return false
	}

	if root != nil && len(arr) == 0 {
		return false
	}
	if root == nil && len(arr) > 0 {
		return false
	}

	return isValidSequence(root.Left, arr[1:]) || isValidSequence(root.Right, arr[1:])

}
