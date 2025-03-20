class Solution:
    def search(self, nums: list[int], target: int) -> int:
        left = 0
        right = len(nums) - 1

        while left <= right:
            mid = (right + left) // 2
            val = nums[mid]
            if val == target:
                return mid
            elif target < val:
                right = mid - 1
            else:
                left = mid + 1
        return -1
