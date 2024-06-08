package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for ; slow != nil && fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {
		if slow == fast {
			break
		}
	}
	if slow != fast {
		return nil
	}
	for node := head; slow != node; node, slow = node.Next, slow.Next {
	}
	return slow
}
