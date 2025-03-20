class Solution:
    def busyStudent(
        self, startTime: list[int], endTime: list[int], queryTime: int
    ) -> int:
        length = len(startTime)
        res = 0
        for i in range(length):
            if startTime[i] <= queryTime <= endTime[i]:
                res += 1

        return res
