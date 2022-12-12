package main

import "fmt"

//
// Definition for singly-linked list.
//
type ListNode struct {
	Val int
	Next *ListNode
}

// print
func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// deleteDuplicates
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
    for curr != nil {
		if prev != nil && curr.Val == prev.Val {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}
	return head
}

// main
func main() {
	ll := &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}
	y := deleteDuplicates(ll)
	y.print()
}