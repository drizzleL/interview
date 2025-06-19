package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
