package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}
	if n > size {
		n %= size
	}
	pre := &ListNode{
		Next: head,
	}
	node := pre
	n = size - n
	for i := 0; i < n; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return pre.Next

}
