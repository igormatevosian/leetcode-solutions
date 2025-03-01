import string


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        counter = {k: 0 for k in string.ascii_lowercase}
        for i in s1:
            counter[i] += 1
        dct = {k: 0 for k in string.ascii_lowercase}
        left = 0
        for i, el in enumerate(s2):
            dct[el] += 1
            if i - left > len(s1) - 1:
                dct[s2[left]] -= 1
                left += 1
            if dct == counter:
                return True
        return False
