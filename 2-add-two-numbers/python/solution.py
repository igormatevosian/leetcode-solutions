from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        res = ListNode()
        current = res
        remains = 0
        while l1 or l2 or remains:
            current.next = ListNode()
            current = current.next

            if l1:
                current.val += l1.val
                l1 = l1.next
            if l2:
                current.val += l2.val
                l2 = l2.next

            current.val += remains
            remains = 1 if current.val > 9 else 0
            current.val %= 10
        return res.next
