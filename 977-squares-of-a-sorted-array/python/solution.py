class Solution:
    def sortedSquares(self, nums: list[int]) -> list[int]:
        res = []
        right = len(nums) - 1
        left = 0

        while left <= right:
            if abs(nums[left]) > abs(nums[right]):
                res.append(nums[left] ** 2)
                left += 1
            else:
                res.append(nums[right] ** 2)
                right -= 1
        return res[::-1]