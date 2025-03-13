class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = 0
        res = 0
        last_seen = {}
        for i, el in enumerate(s):
            if el in last_seen and last_seen[el] >= start:
                start = last_seen[el] + 1
            last_seen[el] = i
            res = max(res, i - start + 1)
        return res