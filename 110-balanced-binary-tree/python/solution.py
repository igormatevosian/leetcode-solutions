from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True

        stack = [(root, False)]
        heights = {}

        while stack:
            node, visited = stack.pop()
            if not node:
                continue

            if visited:
                left_height = heights.get(node.left, 0)
                right_height = heights.get(node.right, 0)

                if abs(left_height - right_height) > 1:
                    return False

                heights[node] = max(left_height, right_height) + 1
            else:
                stack.append((node, True))
                stack.append((node.right, False))
                stack.append((node.left, False))

        return True
