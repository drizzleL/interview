package main

func isPalindrome2(head *ListNode) bool {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}
	l1, l2 := head, head
	for i := 0; i < (size-1)/2; i++ {
		l2 = l2.Next
	}
	l2, l2.Next = l2.Next, nil
	l2 = reverseList(l2)
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
