class Solution:
    def numIslands(self, grid: list[list[str]]) -> int:
        rows, cols = len(grid), len(grid[0])

        def is_island(i, j):
            return 0 <= i < rows and 0 <= j < cols and grid[i][j] == "1"

        def dfs(i, j):
            stack = [(i, j)]
            while stack:
                i, j = stack.pop()
                if is_island(i, j):
                    grid[i][j] = "0"
                    stack.append((i + 1, j))
                    stack.append((i - 1, j))
                    stack.append((i, j + 1))
                    stack.append((i, j - 1))
        
        res = 0
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == "1":
                    res += 1
                    dfs(i, j)
        return res