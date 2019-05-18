func partition(head *ListNode, x int) *ListNode {
    less := &ListNode{}
    greater := &ListNode{}
    
    node := head
    lnode := less
    gnode := greater
    for node != nil {
        if node.Val < x {
            lnode.Next = node
            lnode = node
        } else {
            gnode.Next = node
            gnode = node
        }
        
        node = node.Next
    }
    
    gnode.Next = nil
    lnode.Next = greater.Next
    
    return less.Next
}
