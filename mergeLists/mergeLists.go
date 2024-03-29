package main

import "fmt"

// Definition for a singly linked
// list.
type ListNode struct {
	Val int
	Next *ListNode
}

// print
// Method for ListNode that
// prints out the values.
func (lst *ListNode) print() {
	for lst != nil {
		fmt.Print(lst.Val)
		lst = lst.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		return &ListNode{list1.Val, mergeTwoLists(list1.Next, list2)}
	} else {
		return &ListNode{list2.Val, mergeTwoLists(list1, list2.Next)}
	}

	return nil
}

func main() {
	test()
}

func test() {
	l1 := &ListNode{1, &ListNode{2, nil}}
	l2 := &ListNode{3, &ListNode{5, &ListNode{6, nil}}}
	x := mergeTwoLists(l1, l2)
	x.print()
}