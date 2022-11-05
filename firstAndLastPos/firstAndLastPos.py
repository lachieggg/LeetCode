#!/usr/bin/env python3

class Solution:

    def searchRange(self, nums: list[int], target: int) -> list[int]:
        return [self.findFirstTarget(nums, target), self.findLastTarget(nums, target)]

    def findFirstTarget(self, nums: list[int], target: int) -> list[int]:
        index = 0
        sign = +1
        delta = 1
        length = len(nums)
        while(1):
            if(nums[index] > target and sign > 0):
                sign = -1
                delta = 1
            elif(nums[index] == target and sign < 0):
                delta *= 2
            elif(nums[index] < target and sign < 0):
                sign = +1
                delta = 1
            elif(nums[index] > target and sign > 0):
                delta = 1
                sign = -1

            index += delta*sign
            if(index > length - 1):
                index = length - 1
                sign = -1
                delta = 1
            if(index < 0):
                index = 0
                sign = +1
                delta = 1

            if(nums[index] == target and nums[index-1] != target):
                return index
    
    def findLastTarget(self, nums: list[int], target: int) -> list[int]:
        index = 0
        sign = +1
        delta = 1
        length = len(nums)
        while(1):
            if(nums[index] > target and sign > 0):
                sign = -1
                delta = 1
            elif(nums[index] == target and sign > 0):
                delta *= 2
            elif(nums[index] < target and sign < 0):
                sign = +1
                delta = 1
            elif(nums[index] > target and sign > 0):
                delta = 1
                sign = -1

            index += delta*sign
            if(index > length - 1):
                index = length - 1
                sign = -1
                delta = 1
            if(index < 0):
                index = 0
                sign = +1
                delta = 1

            if(nums[index] == target and nums[index+1] != target):
                return index

s = Solution()

print(s.searchRange([1,1,2,3,3,3,4,5], 3))