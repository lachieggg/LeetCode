#!/usr/bin/env python3

class Solution:
    def subarraySum(self, nums, k):
        instances = 0
        for startIndex in range(0, len(nums), 1):
            for endIndex in range(startIndex, len(nums), 1):
                _array = nums[startIndex: endIndex+1]
                if(sum(_array) == k):
                    instances += 1

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
k = 6

s = Solution()
print(s.subarraySum(nums, k))

