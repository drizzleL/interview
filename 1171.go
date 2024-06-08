package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	dict := map[int]*ListNode{
		0: pre,
	}
	n := head
	var presum int
	for n != nil {
		presum += n.Val
		old := dict[presum]
		if old != nil { // rewind to old, delete between (old,n]
			next := n.Next
			tmpPresum := presum
			for tmp := old.Next; tmp != n; tmp = tmp.Next {
				tmpPresum += tmp.Val
				delete(dict, tmpPresum)
			}

			old.Next = next
			n = next
		} else {
			dict[presum] = n
			n = n.Next
		}
	}

	return pre.Next
}
