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

// Idea:
// Merge two lists, then recursively merge that 
// list with the remaining lists
//

func mergeKLists(lists []*ListNode) *ListNode {
	if(lists == nil) {
		return &ListNode{}
	}

	merged := mergeTwoLists(lists[0], lists[1])
	if len(lists[2:]) == 0 {
		return merged
	} else {
		// No pointer incrementing in Go!
		var l []*ListNode
		for _, z := range lists[2:] {
			l = append(l, z)
		}
		return mergeKLists(append(l, merged))
	}
}

func main() {
	l1 := &ListNode{1, &ListNode{9, nil}}
	l2 := &ListNode{3, &ListNode{5, &ListNode{6, nil}}}
	l3 := &ListNode{4, &ListNode{6, &ListNode{7, nil}}}
	
	lists := []*ListNode{l1, l2, l3}
	x := mergeKLists(lists)
	x.print()
}