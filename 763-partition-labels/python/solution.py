from collections import Counter


class Solution:
    def partitionLabels(self, s: str) -> list[int]:
        cnt = Counter(s)
        letters = set()
        res = []
        start = 0
        remaining = 0

        for i, el in enumerate(s):
            cnt[el] -= 1
            if el not in letters:
                letters.add(el)
                remaining += 1
           
            if cnt[el] == 0:
                remaining -= 1

            if remaining == 0:
                res.append(i - start + 1)
                start = i + 1
                letters = set()

        return res