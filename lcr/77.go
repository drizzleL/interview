package main

func sortList(head *ListNode) *ListNode {
	merge := func(a, b *ListNode) *ListNode {
		pre := &ListNode{}
		node := pre
		for a != nil || b != nil {
			for a != nil && b != nil {
				if a.Val < b.Val {
					node.Next = a
					node = node.Next
					a = a.Next
				} else {
					node.Next = b
					node = node.Next
					b = b.Next
				}
			}
			for a != nil {
				node.Next = a
				node = node.Next
				a = a.Next
			}
			for b != nil {
				node.Next = b
				node = node.Next
				b = b.Next
			}
		}
		return pre.Next
	}
	var sort func(node *ListNode) *ListNode
	sort = func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}
		slow, fast := node, node.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		a, b := node, slow.Next
		slow.Next = nil
		a, b = sort(a), sort(b)
		return merge(a, b)
	}
	return sort(head)
}
