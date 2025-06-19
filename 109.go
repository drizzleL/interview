package main

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	mid := slow.Next
	slow.Next = nil
	ret := &TreeNode{Val: mid.Val}
	ret.Left = sortedListToBST(head)
	ret.Right = sortedListToBST(mid.Next)
	return ret
}
