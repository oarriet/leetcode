package main

/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if isMirror(root.Left, root.Right) {
		return true
	}
	return false
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil && node2 != nil {
		return false
	}
	if node1 != nil && node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left) {
		return true
	}
	return false
}

//func main() {
//	//left
//	lNode4 := TreeNode{
//		Val:   4,
//		Left:  nil,
//		Right: nil,
//	}
//	lNode3 := TreeNode{
//		Val:   3,
//		Left:  nil,
//		Right: nil,
//	}
//	lNode2 := TreeNode{
//		Val:   2,
//		Left:  &lNode3,
//		Right: &lNode4,
//	}
//	//right
//	rNode4 := TreeNode{
//		Val:   4,
//		Left:  nil,
//		Right: nil,
//	}
//	rNode3 := TreeNode{
//		Val:   3,
//		Left:  nil,
//		Right: nil,
//	}
//	rNode2 := TreeNode{
//		Val:   2,
//		Left:  &rNode4,
//		Right: &rNode3,
//	}
//
//	root := TreeNode{
//		Val:   4,
//		Left:  &lNode2,
//		Right: &rNode2,
//	}
//
//	fmt.Printf("Result: %t\n", isSymmetric(&root))
//}
