package main

/*
Given the head of a singly linked list, return true if it is a palindrome.

Example 1:

Input: head = [1,2,2,1]
Output: true

Example 2:

Input: head = [1,2]
Output: false
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func isPalindrome(head *ListNode) bool {
	var fullList []int

	for head != nil {
		fullList = append(fullList, head.Val)
		head = head.Next
	}

	for i := 0; i < len(fullList)/2; i++ {
		if fullList[i] != fullList[len(fullList)-1-i] {
			return false
		}
	}

	return true
}

//func main() {
//	node4 := ListNode{Val: 1, Next: nil}
//	node3 := ListNode{Val: 2, Next: &node4}
//	node2 := ListNode{Val: 2, Next: &node3}
//	node1 := ListNode{Val: 1, Next: &node2}
//
//	println(isPalindrome(&node1))
//
//	//node2 := ListNode{Val: 2, Next: nil}
//	//node1 := ListNode{Val: 1, Next: &node2}
//	//
//	//println(isPalindrome(&node1))
//}
