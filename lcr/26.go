package main

func reorderList(head *ListNode) {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}
	l1, l2 := head, head
	for i := 0; i < (size-1)/2; i++ {
		l2 = l2.Next
	}
	l2, l2.Next = l2.Next, nil
	l2 = reverseList(l2)
	pre := &ListNode{}
	node := pre
	for l1 != nil && l2 != nil {
		node.Next = l1
		l1 = l1.Next
		node = node.Next
		node.Next = l2
		l2 = l2.Next
		node = node.Next
	}
	node.Next = l1
}
