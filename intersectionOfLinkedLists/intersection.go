package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

// print out a linked list
func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// constructLinkedList
// Takes O(n) time in the size of the list
func constructLinkedList(lst []int) *ListNode {
	l := &ListNode{lst[0], nil}
	c := l
	for _, v := range lst[1:] {
		l.Next = &ListNode{v, nil}
		l = l.Next
	}
	return c
}

// reverseLinkedList
// Takes O(n) time in the size of the list
func (l *ListNode) reverseLinkedList() *ListNode {
	if l == nil {
		return nil
	}
	var lst *ListNode
	current := l

	for current != nil {
		lst = &ListNode{current.Val, lst}
		current = current.Next
	}
	
	start := lst

	return start
}

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	lstA := headA.reverseLinkedList()
	lstB := headB.reverseLinkedList()

	// var prev *ListNode

	for lstA.Val == lstB.Val {
		// prev = lstA
		lstA = lstA.Next
		headA = headA.Next
	}

	return headA
}

func main() {
	x := [...]int{1,1,8,4,5}
	y := [...]int{2,3,4,1,8,4,5}
	// 5 4 8 1 4 3 2 
	// 5 4 8 1 1

	lx := constructLinkedList(x[:])
	ly := constructLinkedList(y[:])

	fmt.Println("****")
	z := getIntersectionNode(lx, ly)
	z.print()
}

// We need to find the intersection of two linked lists
// That is, the node after which list A and list B are
// exactly the same
// 
// This means that we need to keep track of
// what previous nodes there are for each, since 
// it is possible that for list A this occurs 
// closer to the start of the list than for list B
//
// eg.
//
// A: 1 -> 2 -> 3 -> 4 -> 2
// B  0 -> 1 -> 2 -> 3 -> 4 -> 2
// 
//
// We should keep a pointer for both A and B
// for where we are currently up to.
//
// These lists are not ordered.
//
// This can be modelled as a recursive problem.
//
// At each point we want to look back and see if 
// the current node for A is in the previous nodes for
// B and vice versa! 
//
// O(1) memory...
//
// Perhaps we can just keep the previous value in memory.
// 
// 
//
// 