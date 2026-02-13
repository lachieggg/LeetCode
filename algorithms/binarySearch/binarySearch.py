#!/usr/bin/env python3

DNE = -1

class Solution:
    def search(self, nums, target):
        print("{}".format(nums))
        if len(nums) == 1:
            if nums[0] == target :
                return 0
            return DNE

        middle = len(nums)//2

        arrayOne = nums[0:middle]
        arrayTwo = nums[middle:]

        if target >= nums[middle]:
            result = self.search(arrayTwo, target)
            if result == DNE:
                return DNE
            return len(arrayOne) + result
        return self.search(arrayOne, target)

if __name__ == "__main__":
    s = Solution()
    
    tests = [
        # (array, target, expected_index)
        ([1], 1, 0),
        ([1], 2, -1),
        ([1, 2], 1, 0),
        ([1, 2], 2, 1),
        ([1, 2], 3, -1),
        ([1, 3, 5], 3, 1),
        ([1, 3, 5], 2, -1),
        ([1, 2, 3, 4], 4, 3),
        ([1, 2, 3, 4], 0, -1),
        ([1, 2, 3, 4, 5], 3, 2),
        ([1, 2, 3, 4, 5], 6, -1),
        ([1, 2, 3, 4, 5, 9], 9, 5),
        ([1, 2, 3, 4, 5, 9, 11], 1, 0),
        ([1, 2, 3, 4, 5, 9, 11], 9, 5),
        ([1, 2, 3, 4, 5, 9, 11], 11, 6),
        ([1, 2, 3, 4, 5, 9, 11], 10, -1),
        ([1, 3, 5, 7, 9, 11, 13, 15], 7, 3),
        ([1, 3, 5, 7, 9, 11, 13, 15], 15, 7),
        ([1, 3, 5, 7, 9, 11, 13, 15], 2, -1),
        ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5, 4),
        ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 10, 9),
        ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 0, -1),
        ([2, 4, 6, 8, 10, 12, 14, 16, 18], 1, -1),
        ([2, 4, 6, 8, 10, 12, 14, 16, 18], 10, 4),
    ]
    
    for arr, target, expected in tests:
        result = s.search(arr, target)
        status = "✅" if result == expected else "❌"
        print(f"{status} array: {str(arr)} target: {target} | expected: {expected} | Got: {result}")
        if result != expected:
            break
        
