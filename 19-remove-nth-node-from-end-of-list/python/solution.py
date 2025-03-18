from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        res = ListNode(0, head)
        slow = res

        for _ in range(n):
            head = head.next

        while head:
            slow = slow.next
            head = head.next

        slow.next = slow.next.next
        return res.next
