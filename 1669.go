package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	pre := &ListNode{Next: list1}
	node := pre
	var apre, bnode *ListNode
	for i := 0; i <= b+1; i++ {
		if i == a {
			apre = node
		}
		if i == b+1 {
			bnode = node
		}
		node = node.Next
	}
	list2End := list2
	for list2End.Next != nil {
		list2End = list2End.Next
	}
	apre.Next = list2
	list2End.Next = bnode.Next
	return pre.Next
}
