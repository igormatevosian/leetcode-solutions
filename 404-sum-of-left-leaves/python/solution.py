from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        res = 0
        stack = [(root, False)]

        while stack:
            node, is_left = stack.pop()
            if not node:
                continue

            if is_left:
                if not node.left and not node.right:
                    res += node.val

            stack.append((node.left, True))
            stack.append((node.right, False))
        return res
