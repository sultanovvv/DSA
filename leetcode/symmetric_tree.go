package leetcode

import "fmt"

/*
	Проверяем поочереди левое и правое поддерево.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SymmetricTree() {
	root := TreeNode{
		Val: 1,
	}
	tn2l := TreeNode{
		Val: 2,
	}
	tn2r := TreeNode{
		Val: 2,
	}
	root.Left = &tn2l
	root.Right = &tn2r

	fmt.Println(calc(root.Left, root.Right))

}

func calc(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return calc(left.Left, right.Right) && calc(left.Right, right.Left)
}
