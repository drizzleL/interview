package main

func findTheWinner(n int, k int) int {
	type node struct {
		num  int
		next *node
	}
	head := &node{
		num: 1,
	}
	no := head
	for i := 2; i <= n; i++ {
		no.next = &node{
			num: i,
		}
		no = no.next
	}
	no.next = head
	pre := no
	for head.next != head {
		for i := 1; i < k; i++ {
			pre = head
			head = head.next
		}
		head = head.next
		pre.next = head
	}
	return head.num
}
