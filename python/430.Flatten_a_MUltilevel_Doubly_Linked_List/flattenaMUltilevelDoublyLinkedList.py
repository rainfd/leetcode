class Solution:
    def flatten(self, head: 'Node') -> 'Node':
        node = head
        while node != None:
            if node.child != None:
                self.addTail(node)
            node = node.next
        return head
            
    def addTail(self, node) -> 'Node':
        if node == None:
            return node
        
        tail = self.getTail(node.child)
        
        tail.next = node.next
        if node.next != None:
            node.next.prev = tail
        
        node.next = node.child
        node.child.prev = node
        
        node.child = None
        
    def getTail(self, node: 'Node') -> 'Node':
        if node == None:
            return None
        
        tail = node
        while tail.next != None:
            tail = tail.next
            
        return tail
