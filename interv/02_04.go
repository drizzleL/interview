package main

func partition(head *ListNode, x int) *ListNode {
	smaller := &ListNode{}
	bigger := &ListNode{}
	n1, n2 := smaller, bigger
	for head != nil {
		if head.Val < x {
			n1.Next = head
			n1 = n1.Next
		} else {
			n2.Next = head
			n2 = n2.Next
		}
		head.Next, head = nil, head.Next
	}
	n1.Next = bigger.Next
	return smaller.Next
}
