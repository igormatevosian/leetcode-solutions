from collections import Counter


class Solution:
    def intersect(self, nums1: list[int], nums2: list[int]) -> list[int]:
        cnt1 = Counter(nums1)
        res = []
        for num in nums2:
            if num in cnt1 and cnt1[num] != 0:
                res.append(num)
                cnt1[num] -= 1
        return res
