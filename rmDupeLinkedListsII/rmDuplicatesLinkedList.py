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
        

        curr = head
        prev = None
        tip = True

        if head.next and head.val == head.next.val:
            tip = False
        
        while(curr):
            if curr.next and curr.val == curr.next.val:
                while(curr and curr.next and curr.val == curr.next.val):
                    curr.next = curr.next.next
                if not prev:
                    prev = curr
                else:
                    prev.next = curr.next
            else:
                prev = curr

            curr = curr.next

        if tip:
            return head
        else:
            return head.next


l = ListNode(2, ListNode(2, ListNode(2, ListNode(3, ListNode(3, ListNode(5, None))))))
s = Solution()
m = s.deleteDuplicates(l)
m.print()