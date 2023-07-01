#!/usr/bin/env python3

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    
    def print(self):
        curr = self
        while(curr):
            print(curr.val)
            curr = curr.next

class Solution:
    def deleteDuplicates(self, head):
        if not head:
            return None

        curr = head.next
        prev = head

        while(curr):
            if curr.val == prev.val:
                # Remove the current node
                prev.next = curr.next
            else:
                prev = curr

            curr = curr.next

        return head


l = ListNode(1, ListNode(1, ListNode(1, ListNode(1, ListNode(1, ListNode(5, None))))))
s = Solution()
m = s.deleteDuplicates(l)
