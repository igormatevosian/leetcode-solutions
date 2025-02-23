class Solution:
    def maxDistToClosest(self, seats: list[int]) -> int:
        start = -1
        res = 0
        for i, el in enumerate(seats):
            if el:
                if start == -1:
                    res = i
                else:
                    res = max(res, (i - start) // 2)
                start = i
        return max(res, len(seats) - 1 - start)