/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 {
        return head
    }
    if head == nil {
        return nil
    }
    
    var length int
    var node, tail, pre, newHead *ListNode
    
    node = head
    i := 1
    for node.Next != nil {
        node = node.Next
        i++
    }
    length = i
    tail = node
    
    if k % length == 0 {
        return head
    }
    
    node = head
    for i := 0; i < (length - k % length) - 1; i++ {
        node = node.Next
    }
    pre = node
    newHead = pre.Next
    tail.Next = head
    pre.Next = nil
    
    return newHead
}
