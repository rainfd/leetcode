func reverseBetween(head *ListNode, m int, n int) *ListNode {
    stack := []*ListNode{}
    
    preHead := &ListNode{ Next: head }
    node := head
    nnode := preHead
    for i := 1; i <= n; i++ {   
        if i < m {
            nnode = node
            node = node.Next
            continue
        } else if i <= n {    
            stack = append(stack, node)
            node = node.Next
        }
    }

    length := len(stack)
    for i := length-1; i >= 0; i-- {
        nnode.Next = stack[i]
        nnode = stack[i]
    }
    nnode.Next = node
    
    return preHead.Next
}
