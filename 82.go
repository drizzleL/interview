package main

func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	node := pre
	for node != nil && node.Next != nil {
		val := node.Next.Val
		if node.Next.Next == nil || node.Next.Next.Val != val { // next exists dup
			node = node.Next
			continue
		}
		for node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		}
	}
	return pre.Next
}
