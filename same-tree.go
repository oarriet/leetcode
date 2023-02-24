package main

/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

//func main() {
//	pnode3 := TreeNode{
//		Val:   3,
//		Left:  nil,
//		Right: nil,
//	}
//	pnode2 := TreeNode{
//		Val:   2,
//		Left:  nil,
//		Right: nil,
//	}
//	pnode1 := TreeNode{
//		Val:   1,
//		Left:  &pnode2,
//		Right: &pnode3,
//	}
//
//	qnode3 := TreeNode{
//		Val:   3,
//		Left:  nil,
//		Right: nil,
//	}
//	qnode2 := TreeNode{
//		Val:   2,
//		Left:  nil,
//		Right: nil,
//	}
//	qnode1 := TreeNode{
//		Val:   1,
//		Left:  &qnode2,
//		Right: &qnode3,
//	}
//	results := isSameTree(&pnode1, &qnode1)
//	fmt.Printf("Results: %v\n", results)
//}
