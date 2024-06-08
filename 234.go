package main

func isPalindrome(head *ListNode) bool {
	equal := func(a, b *ListNode) bool {
		for a != nil && b != nil {
			if a.Val != b.Val {
				return false
			}
			a, b = a.Next, b.Next
		}
		return true
	}
	reverse := func(h *ListNode) *ListNode {
		var last *ListNode
		for h != nil {
			h, last, h.Next = h.Next, h, last
		}
		return last
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var h2 *ListNode
	if fast == nil {
		h2 = slow
	} else {
		h2 = slow.Next
	}
	h2 = reverse(h2)
	return equal(head, h2)
}
