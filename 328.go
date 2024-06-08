package main

func oddEvenList(head *ListNode) *ListNode {
	n := head
	l1pre, l2pre := &ListNode{}, &ListNode{}
	l1, l2 := l1pre, l2pre
	for n != nil {
		l1.Next = n
		l1 = l1.Next
		n = n.Next
		if n != nil {
			l2.Next = n
			l2 = l2.Next
			n = n.Next
		}
	}
	l2.Next = nil
	l1.Next = l2pre.Next
	return l1pre.Next
}
