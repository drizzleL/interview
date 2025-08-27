package main

func getDecimalValue(head *ListNode) int {
	var size int
	for node := head; node != nil; node = node.Next {
		size += 1
	}
	var ret int
	for node := head; node != nil; node = node.Next {
		if node.Val == 1 {
			ret += 1 << (size - 1)
		}
		size -= 1
	}
	return ret
}
