#!/usr/bin/python

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    
    def print(self):
        x = self
        while(x):
            print(x.val)
            x = x.next


def mergeTwoLists(list1, list2):
	if list1 == None:
		return list2

	if list2 == None:
		return list1

	if list1.val < list2.val:
		return ListNode(list1.val, mergeTwoLists(list1.next, list2))
	else:
		return ListNode(list2.val, mergeTwoLists(list1, list2.next))

def main():
	test()

def test():
	l1 = ListNode(1, ListNode(2, None))
	l2 = ListNode(3, ListNode(5, ListNode(6, None)))
	x = mergeTwoLists(l1, l2)
	x.print()

if(__name__ == "__main__"):
    main()