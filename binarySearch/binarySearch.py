#!/usr/bin/env python3

DOES_NOT_EXIST = -1

class Solution:
    def search(self, nums, target):
        #print("Nums {}".format(nums))
        if(len(nums) == 1):
            if(nums[0] == target):
                return 0
            return DOES_NOT_EXIST

        middle = len(nums)//2

        arrayOne = nums[0:middle]
        arrayTwo = nums[middle:]

        #print(arrayOne)
        #print(arrayTwo)

        if(target >= nums[middle]):
            #print("Target {} is greater than {}".format(target, nums[middle]))
            result = self.search(arrayTwo, target)
            if(result == DOES_NOT_EXIST):
                return DOES_NOT_EXIST
            return len(arrayOne) + result
        return self.search(arrayOne, target)

array = [1,2,3,4,5,9,11]
s = Solution()
index = s.search(array, 9)
print(index)