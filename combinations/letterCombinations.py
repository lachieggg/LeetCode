#!/usr/bin/env python3

class Solution:
    def __init__(self):
        self.mapping = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g' ,'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z']
        }

    def letterCombinations(self, digits: str) -> list[str]:
        res = []
        if not digits : return res
        letters = self.mapping.get(digits[0])

        for letter in letters:
            if digits[1:]:
                results = self.letterCombinations(digits[1:])
                for result in results:
                    res.append(letter + result)
            else:
                res.append(letter)

        return res

if __name__ == "__main__":
    s = Solution()
    print(s.letterCombinations("23"))
    print(s.letterCombinations(""))