# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        while head is not None:
            if isinstance(head.val, str):
                return True
            head.val = str(head.val)
            head = head.next
        return False
