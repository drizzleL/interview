package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	l2 = reverse(l2)
	merge(l1, l2)
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	for next != nil {
		next.Next, head, next = head, next, next.Next
	}
	return head
}

func merge(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	node := pre
	for l1 != nil && l2 != nil {
		node.Next = l1
		node = node.Next
		l1 = l1.Next
		node.Next = l2
		node = node.Next
		l2 = l2.Next
	}
	for l1 != nil {
		node.Next = l1
		node = node.Next
		l1 = l1.Next
	}
	for l2 != nil {
		node.Next = l2
		node = node.Next
		l2 = l2.Next
	}
	return node.Next
}
