package main

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dict := map[int]bool{}
	dict[head.Val] = true
	n := head
	for n.Next != nil {
		if dict[n.Next.Val] {
			n.Next = n.Next.Next
			continue
		}
		dict[n.Next.Val] = true
		n = n.Next
	}
	return head
}
