func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for slow != nil && fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            if slow == head {
                return head
            }
            return fast.Next
        }
    }
    return nil
}
