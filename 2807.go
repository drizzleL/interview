package main

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	v1, v2 := head.Val, head.Next.Val
	mid := &ListNode{
		Next: head.Next,
		Val:  gcd(v1, v2),
	}
	head.Next, mid.Next = mid, head.Next
	mid.Next = insertGreatestCommonDivisors(mid.Next)
	return head
}
