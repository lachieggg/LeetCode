
def same(i, j, ss):
    if i - 1 < 0 or i >= len(ss):
        return True

    if j >= len(ss[i]) or ss[i-1][j] != ss[i][j]: 
        return False
   
    return True

class Solution:
    def longestCommonPrefix(self, ss: list[str]) -> str:
        for j in range(len(ss[0])):
            for i in range(len(ss)):
                if not same(i, j, ss):
                    return ss[0][:j]

        return ss[0]

if __name__ == "__main__":
    s = Solution()

    tests = [
        (["", ""], ""),
        (["a", ""], ""),
        (["a", "ab"], "a"),
        (["ab", "a"], "a"),
        (["flower", "flow", "flight"], "fl"),
    ]

    for input_data, expected in tests:
        result = s.longestCommonPrefix(input_data)
        status = "✅" if result == expected else "❌"
        print(f"{status}, input: {input_data}, expected: '{expected}', got: '{result}'")