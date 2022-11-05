#!/usr/bin/env python3

class Solution:
    def __init__(self):
        self.mapping = {
            2: ['a', 'b', 'c'],
            3: ['d', 'e', 'f'],
            4: ['g' ,'h', 'i'],
            5: ['j', 'k', 'l'],
            6: ['m', 'n', 'o'],
            7: ['p', 'q', 'r', 's'],
            8: ['t', 'u', 'v'],
            9: ['w', 'x', 'y', 'z']
        }

    def letterCombinations(self, digits: str) -> list[str]:
        if(len(digits) == 0):
            return []
        letters = self.mapping.get(int(digits[0]))

        res = []
        for letter in letters:
            results = self.letterCombinations(digits[1:])
            if not results:
                res.append(letter)
            else:
                for result in results:
                    res.append(letter + result)
            
        return res

s = Solution()
print(s.letterCombinations("468369592"))
print(s.letterCombinations("23"))
print(s.letterCombinations(""))