from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        stack = [root.left, root.right]

        while stack:
            right = stack.pop()
            left = stack.pop()
            if not right and not left:
                continue
            elif not right or not left:
                return False

            if right.val != left.val:
                return False

            stack.append(left.left)
            stack.append(right.right)

            stack.append(left.right)
            stack.append(right.left)

        return True
