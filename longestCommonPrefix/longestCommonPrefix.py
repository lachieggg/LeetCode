#!/usr/bin/env python3

def same(i, j, strs):
    try:
        if strs[i-1][j] != strs[i][j]: return False
    except IndexError:
        return False
    return True

class Solution:
    def longestCommonPrefix(self, strs: list[str]) -> str:
        if not strs:
            return ""

        strlen = len(strs[0])
        arrlen = len(strs)

        for j in range(strlen):
            for i in range(arrlen):
                if not same(i, j, strs):
                    return strs[0][:j]

        # All the same string
        # Or length == 1    
        return strs[0]

if(__name__ == "__main__"):
    s = Solution()
    arr = ["flower", "flow", "flight"]
    arr = ["flower"]

    print(s.longestCommonPrefix(arr))
    arr = ["dog","racecar","car"]
    print(s.longestCommonPrefix(arr))
    arr = []
    print(s.longestCommonPrefix(arr))