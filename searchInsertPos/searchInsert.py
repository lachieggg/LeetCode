#!/usr/bin/env python3


class Solution(object):
	def searchInsert(self, nums, target):
		"""
		:type nums: List[int]
		:type target: int
		:rtype: int
		"""
		if(len(nums) == 1):
			if(target > nums[0]):
				return 1
			elif(target < nums[0]):
				return 0
			else:
				return 0

		# Split the array in half
		halfLength = len(nums)//2
		arrOne = nums[0:halfLength]
		arrTwo = nums[halfLength:]
		print("Target = {}".format(target))
		print("Orig = {}".format(nums))
		print("Array one = {}".format(arrOne))
		print("Array two = {}".format(arrTwo))

		first = self.searchInsert(arrOne, target)
		second = self.searchInsert(arrTwo, target)
		result = first + second
		return result

	def test(self):
		arr = [1,2,3,4]
		target = 5
		k = self.searchInsert(arr, target)
		print(k)


		arr = [1,3,5,7,9]
		target = 2
		k = self.searchInsert(arr, target)
		print(k)

		arr = [1,4,9,16,25,36]
		target = 10
		k = self.searchInsert(arr, target)
		print(k)


s = Solution()
s.test()
