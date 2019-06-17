func splitListToParts(root *ListNode, k int) []*ListNode {
    length := 0
    for node := root; node != nil; node = node.Next {
        length = length + 1
    }
    
    elements := length / k
    addition := length % k
    
    result := []*ListNode{}
    
    node := root
    for i := 0; i <  k; i++ {
        if node == nil {
            var n *ListNode
            result = append(result, n)
            continue
        }
        
        result = append(result, node)
        
        // raw node 
        elm := elements - 1
        if i < addition {
            elm = elm + 1
        }
        
        for j := 0; j < elm && node != nil ; j++ {
            node = node.Next       
        }
  
        tmp := node
        node = node.Next
        if tmp != nil {
            tmp.Next = nil
        }
    }
    return result
}
