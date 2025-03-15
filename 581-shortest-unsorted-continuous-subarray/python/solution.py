class Solution:
    def findUnsortedSubarray(self, nums: list[int]) -> int:
        min_val, max_val = float("inf"), float("-inf")
        left = len(nums) - 1
        right = 0

        for i in range(len(nums)):
            if nums[i] < max_val:
                right = i
            else:
                max_val = nums[i]

            if nums[-i - 1] > min_val:
                left = len(nums) - i - 1
            else:
                min_val = nums[-i - 1]
                
        if right > left:
            return right - left + 1
        return 0