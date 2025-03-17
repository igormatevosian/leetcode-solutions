class Solution:
    def maxSubArray(self, nums: list[int]) -> int:
        res = float("-inf")
        prefix_sum = 0

        for num in nums:
            prefix_sum += num
            res = max(res, prefix_sum)
            if prefix_sum < 0:
                prefix_sum = 0

        return res
