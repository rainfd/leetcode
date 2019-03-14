# Definition for singly-linked list with a random pointer.
# class RandomListNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.next = None
#         self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        if not head:
            return None
        
        # A->B->C
        # A->A'->B->B'->C->C'
        # double
        node = head
        while node:
            new = RandomListNode(node.label)
            new.next = node.next
            node.next = new
            node = new.next
        
        # copy random pointer
        node = head
        while node:
            if node.random:
                node.next.random = node.random.next
            node = node.next.next
        
        # split
        node = head
        copy_head = head.next
        while node:
            copy_node = node.next
            node.next = node.next.next
            
            if node.next:
                copy_node.next = node.next.next

            node = node.next
            
        return copy_head
