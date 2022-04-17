#!/usr/bin/env python3


class Solution:
    def optimize(self, nums):
        # Remove all duplicates of a particular number
        prevs = []
        for num in nums:
            if(num in prevs):
                continue
            prevs.append(num)
        return prevs

    def threeSum(self, nums):
        nums = self.optimize(nums)
        self.arr = nums
        output = {}

        for i,first in enumerate(nums):
            for j,second in enumerate(nums):
                if(i==j):
                    continue
                pair = ((i, first), (j, second))
                output[pair] = first + second
        
        solutions = []
        for k,third in enumerate(nums):
            for key in output.keys():
                if(output[key] == -third and not k == key[0][0] and not k == key[1][0]):
                    m = [key[0][1], key[1][1], third]
                    m = sorted(m)
                    if(m in solutions):
                        continue
                    solutions.append(sorted(m))

        return solutions

arr = [3,0,-2,-1,1,2,2]
s = Solution().threeSum(arr)
print(s)


s = Solution().optimize(arr)
print(s)