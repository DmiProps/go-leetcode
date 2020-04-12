package problems

import "fmt"

// Problem19 - 19. Remove Nth Node From End of List
func Problem19() {

	head := &ListNode1{Val: 1}
	/*head.Next = &ListNode1{Val: 2}
	head.Next.Next = &ListNode1{Val: 3}
	head.Next.Next.Next = &ListNode1{Val: 4}
	head.Next.Next.Next.Next = &ListNode1{Val: 5}*/

	h := removeNthFromEnd(head, 1)

	for ; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}

}

// ListNode1 - struct
type ListNode1 struct {
	Val  int
	Next *ListNode1
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode1, n int) *ListNode1 {

	a := []*ListNode1{head}
	for h := head.Next; h != nil; h = h.Next {
		a = append(a, h)
	}
	if n == len(a) {
		if n == 1 {
			return nil
		}
		return a[1]
	}
	if n == 1 {
		a[len(a)-2].Next = nil
	} else {
		a[len(a)-n-1].Next = a[len(a)-n+1]
	}

	return head
}
