#!/usr/bin/env python3

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        nodes = []
        curr = head
        if(not curr):
            return False

        while(curr.next and curr.next not in nodes):
            nodes.append(curr)
            curr = curr.next
       
        
        if(curr.next):
            return True
        return False