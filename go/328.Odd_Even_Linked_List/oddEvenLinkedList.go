/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
   if head == nil || head.Next == nil {
       return head
   }
   even_head := head.Next

   node1 := head
   node2 := head.Next
   tail := node1

   for node1 != nil && node2 != nil && node2.Next != nil {
       tail = node1
       node1.Next = node2.Next
       node2.Next= node2.Next.Next
       node1 = node1.Next
       node2 = node2.Next
   }

   if node1 != nil {
       tail = node1
   }
   tail.Next = even_head

   return head
}
