package problems

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Problem876 - 876. Middle of the Linked List
func Problem876() {

	test1 := new(ListNode)
	test1.Val = 1
	test1.Next = new(ListNode)
	test1.Next.Val = 2
	test1.Next.Next = new(ListNode)
	test1.Next.Next.Val = 3
	test1.Next.Next.Next = new(ListNode)
	test1.Next.Next.Next.Val = 4
	test1.Next.Next.Next.Next = new(ListNode)
	test1.Next.Next.Next.Next.Val = 5

	fmt.Println(middleNode(test1).Val)

	test1.Next.Next.Next.Next.Next = new(ListNode)
	test1.Next.Next.Next.Next.Next.Val = 6

	fmt.Println(middleNode(test1).Val)

}

func middleNode(head *ListNode) *ListNode {

	ans := head
	cnt := 1
	for head.Next != nil {
		head = head.Next
		cnt++
		if cnt%2 == 0 {
			ans = ans.Next
		}
	}

	return ans

}
