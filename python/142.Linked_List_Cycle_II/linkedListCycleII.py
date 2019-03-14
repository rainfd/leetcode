# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def detectCycle(self, head):
    	while head is not None:
    		if isinstance(head.val, str):
    			return head
    		head.val = str(head.val)
    		head = head.next
    	return None
    def detectCycle2(self,head):
    	tag = {}
    	while head is not None:
    		tag[head] = 1
    		head = head.next
    		if head in tag:
    			return head
    	return None
