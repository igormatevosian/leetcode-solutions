class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        dct = {}
        for i, el in enumerate(nums):
            if target - el in dct:
                return i, dct[target - el]
            dct[el] = i
        return -1, -1