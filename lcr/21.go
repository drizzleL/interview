package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}
	m := (size - (n % size)) % size
	dummy := &ListNode{Next: head}
	node := dummy
	for i := 0; i < m; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return dummy.Next
}
