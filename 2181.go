package main

func mergeNodes(head *ListNode) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	last := pre
	var flag bool
	for n := head; n != nil; {
		switch {
		case n.Val == 0:
			flag = true
		case flag:
			last.Next = n
			last = n
			flag = false
		default:
			last.Val += n.Val
		}
		n = n.Next
		last.Next = nil
	}
	return pre.Next
}
