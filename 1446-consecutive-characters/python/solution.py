class Solution:
    def maxPower(self, s: str) -> int:
        res = 1
        start = 0
        for i, char in enumerate(s):
            if s[start] != char:
                start = i
            res = max(res, i - start + 1)
        return res
