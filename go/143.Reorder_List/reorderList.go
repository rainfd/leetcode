func reorderList(head *ListNode)  {
    second := splitList(head)
    second = reverseList(second)
    mergeList(head, second)
}

func lenList(head *ListNode) int {
    length := 0
    node := head
    for node != nil {
        node = node.Next
        length = length + 1
    }
    return length
}

func splitList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    length := lenList(head)
    len2 := length / 2
    len1 := length - len2
    
    node := head
    for i := 1; i < len1; i++ {
        node = node.Next
    }
    tmp := node.Next
    node.Next = nil
    
    return tmp
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    node := head
    var preHead *ListNode
    for node != nil {
        tmp := node.Next
        node.Next = preHead
        preHead = node
        node = tmp
    }
    
    return preHead
}

func mergeList(head1, head2 *ListNode) {
    node1, node2 := head1, head2
    for node1 != nil && node2 != nil {
        tmp := node1.Next
        tmp2 := node2.Next
        
        node1.Next = node2
        node2.Next = tmp

        node1 = tmp
        node2 = tmp2
    }
}
