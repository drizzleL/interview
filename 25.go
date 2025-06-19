package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	reverse := func(head *ListNode) (h, t *ListNode) {
		var pre *ListNode
		for node := head; node != nil; {
			next := node.Next
			node.Next = pre
			node, pre = next, node
		}
		return pre, head
	}
	pre := &ListNode{Next: head}
	h1 := pre
	for h1.Next != nil {
		node := h1
		var cnt int
		for i := 0; node.Next != nil && i < k; i++ {
			node = node.Next
			cnt += 1
		}
		next := node.Next
		if cnt < k {
			break
		}
		node.Next = nil
		h2, t2 := reverse(h1.Next)
		h1.Next = h2
		t2.Next = next
		h1 = t2
	}
	return pre.Next
}
