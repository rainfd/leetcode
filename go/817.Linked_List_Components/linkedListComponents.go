/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
    record := &ListNode{}
    
    node := head
    rec := record
    // mark value in record
    // 0, 1, 2, 3, 4,
    // 0, 1, -1,1, 1,
    for node != nil {
        tmp := &ListNode{ Val: -1 }
        for _, n := range G {
            if n == node.Val {
                tmp.Val = n
                break
            }
        }
        rec.Next = tmp
        rec = rec.Next
        node = node.Next
    }
    
    numbers := 0
    rec = record.Next
    previous := -1
    
    for rec != nil {
        if rec.Val != -1 && previous == -1 {
            numbers++
        }
        previous = rec.Val
        rec = rec.Next
    }
    
    return numbers
}
