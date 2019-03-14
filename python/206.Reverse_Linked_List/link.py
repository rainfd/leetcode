class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        h = ListNode(None)
        while head is not None:
            node = ListNode(head.val)
            node.next = h.next
            h.next = node
            head = head.next
        return h.next
