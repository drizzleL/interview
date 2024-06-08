package main

func insertionSortList(head *ListNode) *ListNode {
	pre := &ListNode{
		Val: -5001,
	}
	now := head
	for now != nil {
		next := now.Next
		insert := pre
		for insert.Next != nil && insert.Next.Val < now.Val {
			insert = insert.Next
		}
		now.Next = insert.Next
		insert.Next = now
		now = next
	}
	return pre.Next
}
