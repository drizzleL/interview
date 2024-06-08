package main

func doubleIt(head *ListNode) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var last *ListNode
		for node != nil {
			next := node.Next
			node.Next = last
			last = node
			node = next
		}
		return last
	}
	head = reverse(head)
	var carry int
	node := head
	for ; node.Next != nil; node = node.Next {
		val := node.Val*2 + carry
		node.Val = val % 10
		carry = val / 10
	}
	val := node.Val*2 + carry
	node.Val = val % 10
	carry = val / 10
	if carry != 0 {
		node.Next = &ListNode{
			Val: 1,
		}
	}
	return reverse(head)
}
