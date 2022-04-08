#!/usr/bin/env python3

RECURSION_LIMIT = 10**9

import sys

sys.setrecursionlimit(RECURSION_LIMIT)

class Solution:
    def getSum(self, nums, startIndex, endIndex):
     
        if(startIndex == endIndex):
         
            return nums[startIndex]

        # Check to see if we have a result already to save time
        s = self.sums.get((startIndex, endIndex))
        if(s):
            return s

        s = nums[startIndex] + self.getSum(nums, startIndex+1, endIndex)

        # Store result for later on 
        self.sums[(startIndex, endIndex)] = s
        return s
    
    def allSameSign(self, nums):
        """Check if the rest of the numbers are all the same sign, which indicates no further solutions"""
        if(nums[0] > 0):
            for n in nums:
                if(n < 0):
                    return False
            return True
        elif(nums[0] < 0):
            for n in nums:
                if(n > 0):
                    return False
            return True
        return False
    
    def discardData(self, nums, startIndex):
        """Discard irrelevant data from dictionary once we are past the index to save memory"""
        keys = []
        for key in self.sums.keys():
            if(key[0] <= startIndex):
                keys.append(key)
        
        for key in keys:
            self.sums.pop(key)

    def subarraySum(self, nums, k):
        # Map a tuple of indexes to a sum
        self.sums = {} 
        instances = 0

        for startIndex in range(0, len(nums), 1):
            for endIndex in range(startIndex, len(nums), 1):
                if(self.getSum(nums, startIndex, endIndex) == k):
                    instances += 1
                    if(self.allSameSign(nums[startIndex:])):
                        break
            self.discardData(nums, startIndex)
        return instances


f = open('input', 'r')
lines = f.readlines()
nums = eval(''.join(lines))

# Idea:
#
# Use Dynamic Programming, we can work out the sum of a subset of 
# a subarray first, and then use that to construct the sum faster
#
# I.e. work out the sum of a sub-sub-array and then find out what
# the sum of a sub-array is 
#
#
k = -93

#nums = [1,3,2,4,5,6]
#nums = [0,0]


s = Solution()
print(s.subarraySum(nums, k))

