package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head.Next != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	head.Next = pre
	return head
}
