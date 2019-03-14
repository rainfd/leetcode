/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    if hasCycle
}

func hasCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && slow != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        
        if slow == fast {
            return fast
        }
    }
    return nil
}

func linkedListLength(head *ListNode) int {
    length := 0
    //for node := head; node != nil; node = node.Next {
        //length++
    //}
    node := head
    for node != nil {
        length++
        node = node.Next
    }
    return length
}

func cycleLength(head *ListNode) int {
    if head == nil {
        return 0
    }
    length := 1
    for node := head.Next; node != head; node = node.Next {
        length++
    }
    return length
}
