package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1 = mergeTwoLists(list2, list1)
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
	}

	return list1
}