func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        if fast.Next == nil { break }
        fast = fast.Next.Next
        if slow == fast { break }
    }
    // no cycle
    if fast == nil || fast.Next == nil {
        return nil
    }
    
    slow, fast = head, slow
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}
