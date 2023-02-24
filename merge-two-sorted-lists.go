package main

import "fmt"

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var mergedListHead *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mergedListHead = appendListNode(mergedListHead, &ListNode{list1.Val, nil})
			list1 = list1.Next
		} else {
			mergedListHead = appendListNode(mergedListHead, &ListNode{list2.Val, nil})
			list2 = list2.Next
		}
	}

	if list1 != nil {
		return appendListNode(mergedListHead, list1)
	}
	if list2 != nil {
		return appendListNode(mergedListHead, list2)
	}

	return mergedListHead
}

func appendListNode(list *ListNode, item *ListNode) (head *ListNode) {
	head = list

	if list == nil {
		return item
	}
	for list.Next != nil {
		list = list.Next
	}

	list.Next = item
	return head
}

//func main() {
//	/*
//		Input: list1 = [1,2,4], list2 = [1,3,4]
//		Output: [1,1,2,3,4,4]
//	*/
//	list14 := ListNode{Val: 4, Next: nil}
//	list12 := ListNode{Val: 2, Next: &list14}
//	list11 := ListNode{Val: 1, Next: &list12}
//
//	list24 := ListNode{Val: 4, Next: nil}
//	list23 := ListNode{Val: 3, Next: &list24}
//	list21 := ListNode{Val: 1, Next: &list23}
//
//	listMerged := mergeTwoLists(&list11, &list21)
//
//	printList(listMerged)
//}

func printList(list *ListNode) {
	fmt.Println("Printing List")
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Printf("Done")
}
