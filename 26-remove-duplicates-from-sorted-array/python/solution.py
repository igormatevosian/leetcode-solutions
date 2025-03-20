class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        if len(nums) == 1:
            return 1

        index = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[index]:
                index += 1
                nums[index] = nums[i]

        return index + 1
