package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{scenario: " Successful scenario. Merge Big lists", function: bigLinkedListTest},
		{scenario: " Successful scenario. Merge Nil lists", function: nilLinkedListTest},
		{scenario: " Successful scenario. Merge Single lists", function: singleElemLinkedListTest},
		{scenario: " Successful scenario. Merge Half-Nil lists", function: halfNilLinkedListTest},
	}

	for _, test := range tests {
		testFunc := test.function
		t.Run(test.scenario, func(t *testing.T) {
			t.Parallel()
			testFunc(t)
		})
	}
}

func bigLinkedListTest(t *testing.T) {
	al1 := ListNode{Val: 1}
	al2 := ListNode{Val: 2}
	al3 := ListNode{Val: 5}
	al4 := ListNode{Val: 8}

	bl1 := ListNode{Val: 3}
	bl2 := ListNode{Val: 4}
	bl3 := ListNode{Val: 6}
	bl4 := ListNode{Val: 7}

	al1.Next = &al2
	al2.Next = &al3
	al3.Next = &al4

	bl1.Next = &bl2
	bl2.Next = &bl3
	bl3.Next = &bl4

	rl1 := al1
	rl2 := al2
	rl3 := bl1
	rl4 := bl2

	rl5 := al3
	rl6 := bl3
	rl7 := bl4
	rl8 := al4

	rl1.Next = &rl2
	rl2.Next = &rl3
	rl3.Next = &rl4

	rl4.Next = &rl5
	rl5.Next = &rl6
	rl6.Next = &rl7
	rl7.Next = &rl8

	res := mergeTwoLists(&al1, &bl1)

	if !reflect.DeepEqual(*res, rl1) {
		t.FailNow()
	}
}

func nilLinkedListTest(t *testing.T) {
	res := mergeTwoLists(nil, nil)

	if res != nil {
		t.FailNow()
	}
}

func singleElemLinkedListTest(t *testing.T) {
	el1 := ListNode{Val: 10}
	fl1 := ListNode{Val: 1}
	res := mergeTwoLists(&el1, &fl1)

	if !reflect.DeepEqual(*res, ListNode{1, &ListNode{10, nil}}) {
		t.FailNow()
	}

}

func halfNilLinkedListTest(t *testing.T) {
	fl1 := ListNode{Val: 1}
	res := mergeTwoLists(nil, &fl1)

	if !reflect.DeepEqual(*res, ListNode{1, nil}) {
		t.FailNow()
	}

}
