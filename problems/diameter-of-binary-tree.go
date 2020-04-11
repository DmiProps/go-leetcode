package problems

import "fmt"

// TreeNode - struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Problem543 - 543. Diameter of Binary Tree
func Problem543() {

	//[-64,12,18,-4,-53,null,76,null,-51,null,null,-93,3,null,-31,47,null,3,53,-81,33,4,null,-51,-44,-60,11,null,null,null,null,78,null,-35,-64,26,-81,-31,27,60,74,null,null,8,-38,47,12,-24,null,-59,-49,-11,-51,67,null,null,null,null,null,null,null,-67,null,-37,-19,10,-55,72,null,null,null,-70,17,-4,null,null,null,null,null,null,null,3,80,44,-88,-91,null,48,-90,-30,null,null,90,-34,37,null,null,73,-38,-31,-85,-31,-96,null,null,-18,67,34,72,null,-17,-77,null,56,-65,-88,-53,null,null,null,-33,86,null,81,-42,null,null,98,-40,70,-26,24,null,null,null,null,92,72,-27,null,null,null,null,null,null,-67,null,null,null,null,null,null,null,-54,-66,-36,null,-72,null,null,43,null,null,null,-92,-1,-98,null,null,null,null,null,null,null,39,-84,null,null,null,null,null,null,null,null,null,null,null,null,null,-93,null,null,null,98]

	test1 := &TreeNode{Val: 1}
	test1.Left = &TreeNode{Val: 2}
	test1.Right = &TreeNode{Val: 3}
	test1.Left.Left = &TreeNode{Val: 4}
	test1.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(test1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}
	ans, _ := calc(root)

	return ans

}

func calc(n *TreeNode) (int, int) {

	dLeft := 0
	dRight := 0
	hLeft := 0
	hRight := 0
	d := 0
	h := 0
	if n.Left != nil {
		dLeft, hLeft = calc(n.Left)

	}
	if n.Right != nil {
		dRight, hRight = calc(n.Right)
	}
	if dLeft > dRight {
		d = dLeft
	} else {
		d = dRight
	}
	if hLeft > hRight {
		h = hLeft
	} else {
		h = hRight
	}
	if hLeft+hRight > d {
		d = hLeft + hRight
	}
	return d, h + 1

}
