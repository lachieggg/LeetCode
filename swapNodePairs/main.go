package main

import (
	"fmt"
)

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs keeps track of the head
// and calls the 'sp' function.
func swapPairs(head *ListNode) *ListNode {
	sp(head)
	return head
}

// sp performs the work of swapping
// listnode values, recursively, until
// it hits the end of the linked list.
func sp(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	swap(head, head.Next)
	head = head.Next.Next
	return sp(head)
}

func swap(n1 *ListNode, n2 *ListNode) {
	x := n1.Val
	n1.Val = n2.Val
	n2.Val = x
}

func main() {
	x := &ListNode{1, nil}
	swapPairs(x)
	printLinkedList(x)
	x = &ListNode{}
	swapPairs(x)
	printLinkedList(x)
	x = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	swapPairs(x)
	printLinkedList(x)
	x = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	swapPairs(x)
	printLinkedList(x)
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println("\n------------------------")
}
