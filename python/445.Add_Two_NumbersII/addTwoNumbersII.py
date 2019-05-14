# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        node1, node2 = l1, l2
        n1, n2 = 0, 0
        while node1 != None:
            n1 = n1 * 10 + node1.val
            node1 = node1.next
        while node2 != None:
            n2 = n2 * 10 + node2.val
            node2 = node2.next
            
        # Python Big Number support
        n3 = n1 + n2
        
        # for num to list bit by bit
        n3_list = list(str(n3))
        
        # for list to linked list
        head = ListNode(0)
        node = head
        for _, n in enumerate(n3_list):
            tmp = ListNode(n)
            node.next = tmp
            node = node.next
            
        return head.next
