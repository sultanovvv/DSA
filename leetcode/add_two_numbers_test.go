package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l3 := &ListNode{Val: 3, Next: nil}
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: l3}}
	rootTest := reverseList(l1)
	fmt.Printf("%+v", rootTest)
}
