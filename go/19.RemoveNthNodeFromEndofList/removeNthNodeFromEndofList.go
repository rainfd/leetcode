package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	sign := head
	for i := 0; i < n; i++ {
		sign = sign.Next
	}

	preHead := &ListNode{Next: head}
	pre := preHead
	for sign != nil {
		pre = pre.Next
		sign = sign.Next
	}
	pre.Next = pre.Next.Next

	return preHead.Next
}
