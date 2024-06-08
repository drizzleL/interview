package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverseList(l1), reverseList(l2)
	var carry int
	head := &ListNode{}
	node := head
	for l1 != nil || l2 != nil {
		for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
			sum := l1.Val + l2.Val + carry
			sum, carry = sum%10, sum/10
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
		for ; l1 != nil; l1 = l1.Next {
			sum := l1.Val + carry
			sum, carry = sum%10, sum/10
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
		for ; l2 != nil; l2 = l2.Next {
			sum := l2.Val + carry
			sum, carry = sum%10, sum/10
			node.Next = &ListNode{
				Val: sum,
			}
			node = node.Next
		}
	}
	if carry == 1 {
		node.Next = &ListNode{Val: 1}
		node = node.Next
	}
	return reverseList(head.Next)
}
