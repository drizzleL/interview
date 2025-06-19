package main

func swapNodes(head *ListNode, k int) *ListNode {
	var size int
	for node := head; node != nil; node = node.Next {
		size += 1
	}
	k = (k - 1) % size
	findKth := func(head *ListNode, k int) *ListNode {
		for i := 0; i < k; i++ {
			head = head.Next
		}
		return head
	}
	a, b := findKth(head, k), findKth(head, size-k-1)
	a.Val, b.Val = b.Val, a.Val
	return head
}
