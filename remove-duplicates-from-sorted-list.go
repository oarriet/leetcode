package main

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	} else {
		deleteDuplicates(head.Next)
	}

	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}

	return head
}

//func main() {
//	nodeC := ListNode{Val: 2, Next: nil}
//	nodeB := ListNode{Val: 1, Next: &nodeC}
//	nodeA := ListNode{Val: 1, Next: &nodeB}
//	resultingList := deleteDuplicates(&nodeA)
//	printList(resultingList)
//
//	nodeE := ListNode{Val: 3, Next: nil}
//	nodeD := ListNode{Val: 3, Next: &nodeE}
//	nodeC = ListNode{Val: 2, Next: &nodeD}
//	nodeB = ListNode{Val: 1, Next: &nodeC}
//	nodeA = ListNode{Val: 1, Next: &nodeB}
//	resultingList = deleteDuplicates(&nodeA)
//	printList(resultingList)
//
//	nodeC = ListNode{Val: 1, Next: nil}
//	nodeB = ListNode{Val: 1, Next: &nodeC}
//	nodeA = ListNode{Val: 1, Next: &nodeB}
//	resultingList = deleteDuplicates(&nodeA)
//	printList(resultingList)
//}
