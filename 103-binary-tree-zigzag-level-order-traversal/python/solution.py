from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> list[list[int]]:
        res = []
        stack = [root]
        left_to_right = True
        while stack:
            new_stack = []
            level_res = []
            while stack:
                node = stack.pop()
                if not node:
                    continue

                level_res.append(node.val)
                new_stack.append(node.left)
                new_stack.append(node.right)

            if level_res:
                if left_to_right:
                    res.append(level_res)
                else:
                    res.append(level_res[::-1])
            left_to_right = not left_to_right
            stack = new_stack[::-1]

        return res
