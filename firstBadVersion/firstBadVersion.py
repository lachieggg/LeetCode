#!/usr/bin/env python3

# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        index = 0
        sign = +1
        delta = 1
        while(1):
            if(not isBadVersion(index) and sign > 0):
                delta *= 2
            elif(isBadVersion(index) and sign < 0):
                delta *= 2
            elif(not isBadVersion(index) and sign < 0):
                delta = 1
                sign = 1
            elif(isBadVersion(index) and sign > 0):
                delta = 1
                sign = -1

            index += delta*sign
            if(not isBadVersion(index-1) and isBadVersion(index)):
                return index