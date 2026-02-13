#!/usr/bin/env python3

# This problem decomposes into a Fibonacci Sequence.

class Solution(object):
	def __init__(self):
		self.results = {}

	def initialize(self, n):
		for x in range(1,n-1):
			print(x)
			self.results[x] = self.climbStairs(x)
			print(self.results)

	def climbStairs(self, n):
		"""
		:type n: int
		:rtype: int
		"""
		if(self.results == {}):
			self.initialize(n)

		if(n == 1):
			return 1
		if(n == 2):
			return 2

		if(n in self.results.keys()):
			return self.results.get(n)

		return self.climbStairs(n-1) + self.climbStairs(n-2)


if __name__ == "__main__":
	s = Solution()
	s.initialize(45)
	print(s.climbStairs(45))
