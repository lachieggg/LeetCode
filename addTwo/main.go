package main

import (
	"fmt"
)

// Leetcode
// Add Two Numbers

const (
	debug = false
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	var ln ListNode = ListNode{}
	var carry int
	var sum int

	if l1 == nil {
		sum = l2.Val
		calculateCarry(&sum, &carry)

		ln.Val = sum
		result, shouldReturn := executeCarry(carry, nil, l2.Next, ln)
		if shouldReturn {
			return result
		}
		ln.Next = addTwoNumbers(nil, l2.Next)
		return &ln
	}

	if l2 == nil {
		sum = l1.Val
		calculateCarry(&sum, &carry)

		ln.Val = sum
		result, shouldReturn := executeCarry(carry, l1.Next, nil, ln)
		if shouldReturn {
			return result
		}
		ln.Next = addTwoNumbers(nil, l1.Next)
		return &ln
	}

	sum = l1.Val + l2.Val
	// check if need to carry
	calculateCarry(&sum, &carry)

	ln.Val = sum
	result, shouldReturn := executeCarry(carry, l1.Next, l2.Next, ln)
	if shouldReturn {
		return result
	}

	ln.Next = addTwoNumbers(l1.Next, l2.Next)

	return &ln
}

// calculateCarry checks whether we need to carry
// and if so, reduces the current sum by 10
// and sets the carry
func calculateCarry(sum *int, carry *int) {
	if *sum <= 9 {
		return
	}
	*carry = 1
	*sum -= 10
}

// executeCarry
// actually adds the carried integer to the next column
// if we are are the end of the Linked List, then carry
// to a new column
func executeCarry(carry int, l1Next *ListNode, l2Next *ListNode, ln ListNode) (*ListNode, bool) {
	if carry <= 0 {
		return nil, false
	}
	// if both nil, carry into a new column
	if l1Next == nil && l2Next == nil {
		ln.Next = &ListNode{carry, nil}
		return &ln, true
	}

	// otherwise, add to whichever is non nil
	if l1Next == nil {
		l2Next.Val = l2Next.Val + carry
	} else {
		l1Next.Val = l1Next.Val + carry
	}

	return &ln, false
}

func printLinkedList(l *ListNode) {
	fmt.Printf("%v ", l.Val)
	if l.Next == nil {
		fmt.Println()
		return
	}
	printLinkedList(l.Next)
}

func makeLinkedList(nums []int) *ListNode {
	var head *ListNode = &ListNode{}
	l := head
	length := len(nums)
	for i, n := range nums {
		l.Val = n
		if i < length-1 {
			l.Next = &ListNode{}
		}
		l = l.Next
	}
	return head
}

func main() {
	// [7 0 8]
	l1 := makeLinkedList([]int{2, 4, 3})
	l2 := makeLinkedList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	printLinkedList(result)

	// [8,9,9,9,0,0,0,1]
	l1 = makeLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = makeLinkedList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	printLinkedList(result)

	return
}
