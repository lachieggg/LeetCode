#!/usr/bin/env python

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head):
        nodes = []
        curr = head
        if(not curr):
            return
        
        while(curr.next and curr.next not in nodes):
            nodes.append(curr)
            curr = curr.next
       
        if(curr.next):
            # There is a cycle
            return nodes.index(curr.next)
        
        return


if __name__ == "__main__":
    s = Solution()
    h = ListNode(3)
    h.next = ListNode(2)
    h.next.next = ListNode(0)
    h.next.next.next = ListNode(-4)
    h.next.next.next = h.next

    print(s.detectCycle(h))