package main

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	for node := head; node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	return head
}
