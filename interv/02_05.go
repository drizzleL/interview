package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ansPre := &ListNode{}
	n := ansPre
	val := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		n.Next = &ListNode{Val: val % 10}
		n = n.Next
		val /= 10
	}
	if val != 0 {
		n.Next = &ListNode{Val: 1}
	}
	return ansPre.Next
}
