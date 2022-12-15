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

// reverseList
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var lst *ListNode
	current := head

	for current != nil {
		lst = &ListNode{current.Val, lst}
		current = current.Next
	}
	
	start := lst

	return start
}

func main() {
	x := [...]int{2,3,4,1,8,4,5}

	lx := constructLinkedList(x[:])

	y := reverseList(lx)
	y.print()
}