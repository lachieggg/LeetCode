#!/usr/bin/env python3

class Solution(object):
    def isPalindrome(self, s):
        for i in range(len(s)-1):
            if(s[i] != s[-(i+1)]):
                return False
        
        return True
    
    def futureStrings(self, s):
        """Return a set of substrings of s that begin with the first letter of s"""
        substrings = []
        length = len(s)
        for i in range(1, length+1):
            substrings.append(s[0:i])
        
        return substrings
    
    def allSubstrings(self, s):
        """Returns all substrings of s"""
        if(len(s) == 0):
            return []

        substrings = []
        for i in range(len(s)):
            for substring in self.futureStrings(s[i:]):
                substrings.append(substring)
        
        return substrings


    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        palindromes = []
        substrings = self.allSubstrings(s)
        for substring in substrings:
            if(self.isPalindrome(substring)):
                palindromes.append(substring)
        
        if(not palindromes):
            return None

        return max(palindromes, key=len)



s = Solution()
substrings = s.longestPalindrome('hello')
print(substrings)