/*
206. Reverse Linked List
Easy
Topics
Companies
Given the head of a singly linked list, reverse the list, and return the reversed list.



Example 1:


Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:


Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000


Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	return prev

}

func main() {
	/*
		Input: list1 = [1,2,3,4,5]
		Output: [5,4,3,2,1]
	*/
	list15 := ListNode{Val: 5, Next: nil}
	list14 := ListNode{Val: 4, Next: &list15}
	list13 := ListNode{Val: 3, Next: &list14}
	list12 := ListNode{Val: 2, Next: &list13}
	list11 := ListNode{Val: 1, Next: &list12}

	listMerged := reverseList(&list11)

	printList(listMerged)
}

func printList(list *ListNode) {
	fmt.Println("Printing List")
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Printf("Done")
}
