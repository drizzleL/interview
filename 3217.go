package main

func modifiedList(nums []int, head *ListNode) *ListNode {
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
	}
	pre := &ListNode{Next: head}
	for n := pre; n.Next != nil; {
		if dict[n.Next.Val] {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return pre.Next
}
