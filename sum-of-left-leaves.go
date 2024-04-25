package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	ret := 0
	if root != nil {
		if root.Left != nil && root.Left.Right == nil && root.Left.Left == nil {
			ret += root.Left.Val
		}
		if root.Left != nil {
			ret += sumOfLeftLeaves(root.Left)
		}
		if root.Right != nil {
			ret += sumOfLeftLeaves(root.Right)
		}
	}
	return ret
}

func main() {
	t := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20,
		Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(sumOfLeftLeaves(t))
}
