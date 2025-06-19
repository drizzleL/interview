package main

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	reverse := func(head *ListNode) (h, t *ListNode) {
		tail := head
		var pre *ListNode
		for head != nil {
			next := head.Next
			head.Next = pre
			pre, head = head, next
		}
		return pre, tail
	}
	pre := &ListNode{
		Next: head,
	}
	for size := 1; pre.Next != nil; size++ {
		h1 := pre.Next
		node := pre
		var cnt int
		for i := 0; node.Next != nil && i < size; i++ {
			node = node.Next
			cnt += 1
		}
		next := node.Next
		if cnt%2 != 0 {
			pre = node
			continue
		}
		node.Next = nil
		h2, t2 := reverse(h1)
		pre.Next = h2
		t2.Next = next
		pre = t2
	}
	return head
}
