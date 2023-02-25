package main

/*

 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lMaxDepth := maxDepth(root.Left)
	rMaxDepth := maxDepth(root.Right)

	if lMaxDepth >= rMaxDepth {
		return lMaxDepth + 1
	} else {
		return rMaxDepth + 1
	}
}

//func main() {
//	node7 := TreeNode{
//		Val:   7,
//		Left:  nil,
//		Right: nil,
//	}
//	node15 := TreeNode{
//		Val:   15,
//		Left:  nil,
//		Right: nil,
//	}
//	node20 := TreeNode{
//		Val:   20,
//		Left:  &node15,
//		Right: &node7,
//	}
//	node9 := TreeNode{
//		Val:   9,
//		Left:  nil,
//		Right: nil,
//	}
//	root := TreeNode{
//		Val:   3,
//		Left:  &node9,
//		Right: &node20,
//	}
//
//	fmt.Printf("Result: %d\n", maxDepth(&root))
//}
