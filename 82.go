package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Next: head, Val: head.Val - 1}
	for pre := preHead; pre != nil; {
		if pre.Next == nil {
			break
		}
		next := pre.Next
		if next.Next == nil {
			break
		}
		if next.Next.Val != next.Val {
			pre = pre.Next
			continue
		}
		val := next.Val
		for next != nil && next.Val == val {
			next = next.Next
		}
		pre.Next = next
	}
	return preHead.Next
}
