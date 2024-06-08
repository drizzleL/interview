package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	getSize := func(head *ListNode) (ret int) {
		for node := head; node != nil; node = node.Next {
			ret += 1
		}
		return ret
	}
	sizeA, sizeB := getSize(headA), getSize(headB)
	if sizeA > sizeB { // a go first
		for i := 0; i < sizeA-sizeB; i++ {
			headA = headA.Next
		}
	} else { // b go first
		for i := 0; i < sizeB-sizeA; i++ {
			headB = headB.Next
		}
	}
	for ; headA != headB; headA, headB = headA.Next, headB.Next {
	}
	return headA
}
