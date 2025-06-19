package main

func splitListToParts(head *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, k)
	var size int
	for n := head; n != nil; n = n.Next {
		size += 1
	}
	avgSize := size / k
	var i int
	for ; i < size-avgSize*k; i++ {
		ret[i] = head
		for j := 0; j < avgSize; j++ {
			head = head.Next
		}
		if head != nil {
			head, head.Next = head.Next, nil
		}
	}
	for ; i < k; i++ {
		ret[i] = head
		for j := 0; j < avgSize-1; j++ {
			head = head.Next
		}
		if head != nil {
			head, head.Next = head.Next, nil
		}
	}
	return ret
}
