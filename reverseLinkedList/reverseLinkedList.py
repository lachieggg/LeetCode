#!/usr/bin/env python3

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def print(self):
        l = self
        while(l):
            print(l.val)
            l = l.next

class Solution:
    def reverseList(self, head):
        if head == None:
            return None

        start = None
        lst = start

        current = head

        while(current):
            if(not current == None):
                lst = ListNode(current.val, lst)
                current = current.next
        
        return lst
    

if __name__ == "__main__":
    l = ListNode(5, ListNode(4, ListNode(3, ListNode(2, ListNode(1, None)))))
    s = Solution()
    m = s.reverseList(l)
    print("\n\n")
    m.print()
