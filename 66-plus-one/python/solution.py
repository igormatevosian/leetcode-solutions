class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        for i in range(len(digits)-1, -1, -1):
            if digits[i] == 9:
                digits[i] = 0
                continue
            digits[i] += 1
            return digits
        digits = [1] + digits
        return digits