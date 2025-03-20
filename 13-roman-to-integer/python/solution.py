class Solution:
    def romanToInt(self, s: str) -> int:
        dct = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

        res = 0
        n = len(s)

        for i in range(n):
            if i < n - 1 and dct[s[i]] < dct[s[i + 1]]:
                res -= dct[s[i]]
            else:
                res += dct[s[i]]

        return res
