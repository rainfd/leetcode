/*
 * 1. find the tail of A
 * 2. link tailA to B head, so link B is a cycle linked list
 * 3. find the cycle point in linked b
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodeA, nodeB := headA, headB
    flagA, flagB := false, false
    for { 
        if nodeA == nodeB {
            return nodeA
        }
        
        if nodeA == nil {
            if flagA {
                break
            } else {
                flagA = true
                nodeA = headB
            }
        } else {
            nodeA = nodeA.Next
        }
        
        if nodeB == nil {
            if flagB {
                break
            } else {
                flagB = true
                nodeB = headA
            }
        } else {
            nodeB = nodeB.Next
        }

    }
    return nil
}


func getIntersectionNodeB(headA, headB *ListNode) *ListNode {
    lengthA := linkedListLength(headA)
    lengthB := linkedListLength(headB)
    delta := lengthA - lengthB
    long, short := headA, headB
    if delta < 0  {
        long, short = short, long
        delta = -delta
    }
    // forward delta step
    for i := 0; i < delta; i++ {
        long = long.Next
    }
    for long != nil && short != nil {
        if long == short {
            return long
        }
        long = long.Next
        short = short.Next
    }
    
    return nil
}

func linkedListLength(head *ListNode) int {
    node := head
    length := 0
    for node != nil {
        length++
        node = node.Next
    }
    return length
}
