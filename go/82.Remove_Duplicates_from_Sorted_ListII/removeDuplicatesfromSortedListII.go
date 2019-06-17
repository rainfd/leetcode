func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	preHead := &ListNode{
		Val:  head.Val - 1,
		Next: head,
	}
	dup := false
	node := head
	pre := preHead
	for node.Next != nil {
		if node.Val == node.Next.Val {
			dup = true
		} else {
			if dup == true {
				pre.Next = node.Next
				dup = false
			} else {
				pre = node
			}
		}
		node = node.Next
	}
	if dup == true {
		pre.Next = nil
	}
	return preHead.Next
}
