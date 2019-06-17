// 876. Middle of the Linked List


func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast, slow = fast.Next.Next, slow.Next
    }
    return slow
}
