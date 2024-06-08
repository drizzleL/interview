package main

func removeNodes(head *ListNode) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var last *ListNode
		for node != nil {
			next := node.Next
			node.Next = last
			last = node
			node = next
		}
		return last
	}
	head = reverse(head)
	for n := head; n.Next != nil; {
		if n.Next.Val < n.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return reverse(head)
}
