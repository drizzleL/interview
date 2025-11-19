package main

func deleteMiddle(head *ListNode) *ListNode {
	var size int
	for n := head; n != nil; n = n.Next {
		size += 1
	}
	if size == 1 {
		return nil
	}
	node := head
	for i := size/2 - 1; i > 0; i-- {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}
