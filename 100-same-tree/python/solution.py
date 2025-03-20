from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        stack = [p, q]

        while stack:
            second = stack.pop()
            first = stack.pop()

            if not second and not first:
                continue
            elif not second or not first:
                return False

            if first.val != second.val:
                return False

            stack.append(first.left)
            stack.append(second.left)
            stack.append(first.right)
            stack.append(second.right)
        return True
