package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	_, root := reverseLinkedList(head)
	return root
}

func reverseLinkedList(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	curr := head
	head, root := reverseLinkedList(head.Next)
	head.Next = curr

	curr.Next = nil

	return curr, root
}
