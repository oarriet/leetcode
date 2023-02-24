package main

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	results = append(results, inorderTraversal(root.Left)...)
	results = append(results, root.Val)
	results = append(results, inorderTraversal(root.Right)...)

	return results
}

//func main() {
//	node3 := TreeNode{
//		Val:   3,
//		Left:  nil,
//		Right: nil,
//	}
//	node2 := TreeNode{
//		Val:   2,
//		Left:  &node3,
//		Right: nil,
//	}
//	node1 := TreeNode{
//		Val:   1,
//		Left:  nil,
//		Right: &node2,
//	}
//	results := inorderTraversal(&node1)
//	fmt.Printf("Results: %v\n", results)
//}
