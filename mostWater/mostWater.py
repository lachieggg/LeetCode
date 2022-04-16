#!/usr/bin/env python3

DEBUG = False

class Solution:
    def maxAreaBruteForce(self, height):
        """Brute force approach, polynomial time algorithm"""
        # End of array
        eoa = len(height)
        # Initialize most water var
        most = 0
        # Iterate
        for start in range(0, eoa, 1):
            for end in range(start+1, eoa, 1):
                width = end - start
                h = min(height[start], height[end])
                water = h * width
                if(water > most):
                    most = water
        
        return most

    def bestLHS(self, height, index):
        """Determine whether a height is the best we can do for LHS"""
        val = height[index]
        for v in height[index:]:
            if(v > val):
                return False
        return True

    def bestRHS(self, height, index, maxLeft):
        """Determine whether a height is the best we can do for RHS"""
        val = height[index]
        for v in height[maxLeft:index]:
            if(v > val):
                return False
        return True


    def area(self, height, start, end):
        """Compute the area for some indexes"""
        width = end - start
        h = min(height[start], height[end])
        area = h * width
        return area
    
    def loadInput(self):
        f = open('input', 'r')
        lines = f.readlines()
        height = eval(''.join(lines))
        return height


    def maxAreaPoly(self, height):
        """Polynomial time solution"""
        left = 0
        length = len(height)

        # Calculate the range of indexes we wish
        # to search through
        for maxLeft in range(0, length, 1):
            if(self.bestLHS(height, maxLeft)):
                break
        for minRight in range(length-1, maxLeft-1, -1):
            if(self.bestRHS(height, minRight, maxLeft)):
                break

        best = self.area(height, maxLeft, minRight)
        for left in range(0, maxLeft+1, 1):
            for right in range(length-1, minRight-1, -1):
                if(self.area(height, left, right) > best):
                    best = self.area(height, left, right)
        return best

    def maxArea(self, height):
        """Linear time solution found with assistance of a friend"""
        left = 0
        right = len(height)-1
        best = self.area(height, left, right)
        while(left < right):
            area = self.area(height, left, right)
            best = area if area > best else best
            if(height[left] > height[right]):
                right -= 1
            else:
                left += 1
        return best

s = Solution()
height = s.loadInput()
answer = s.maxArea(height)
print(answer)