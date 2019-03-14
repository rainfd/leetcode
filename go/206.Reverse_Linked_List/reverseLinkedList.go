/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    node := head
    for node && node.Next {
	tmp := node.Next
        node.Next = tmp.Next
	tmp.Next = head
	head = tmp
    }
    return head
}
