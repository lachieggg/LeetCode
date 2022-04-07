#!/usr/bin/env python3

class Solution:
    def subarraySum(self, nums, k):
        _subArrays = self.subArrays(nums)
        instances = 0
        for _array in _subArrays:
            print(_array)
            if(self.sumOfArray(_array) == k):
                print("True")
                instances += 1
            else:
                print("False")
        
        return instances
        
    def subArrays(self, nums):
        _subArrays = []
        for startIndex in range(0, len(nums), 1):
            for endIndex in range(startIndex, len(nums), 1):
                _subArrays.append(nums[startIndex: endIndex+1])
            
        
        return _subArrays
    
    def sumOfArray(self, _array):
        _sum = 0
        for x in _array:
            _sum += x
        return _sum
    

nums = [1,2,3,4]
k = 6

s = Solution()
print(s.subArrays(nums))
print(s.subarraySum(nums, k))

