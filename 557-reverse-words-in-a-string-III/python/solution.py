class Solution:
    def reverseWords(self, s: str) -> str:
        res = []
        start = 0

        for i, char in enumerate(s):
            if char == " ":
                res.extend(
                    s[i - 1 : start - 1 if start else None : -1] if start != i else []
                )
                res.append(" ")
                start = i + 1

        res.extend(s[len(s) - 1 : start - 1 if start else None : -1])

        return "".join(res)
