package main

import "math"

func nodesBetweenCriticalPoints(head *ListNode) []int {
	var pre *ListNode
	isCritical := func(n *ListNode) bool {
		if pre == nil || n.Next == nil {
			return false
		}
		return (n.Val > pre.Val && n.Val > n.Next.Val) || (n.Val < pre.Val && n.Val < n.Next.Val)
	}
	minDis := math.MaxInt32
	oldIdx, lastIdx := -1, -1
	for i := 0; head != nil; i, head = i+1, head.Next {
		if isCritical(head) {
			if oldIdx == -1 {
				oldIdx = i
			} else {
				minDis = min(minDis, i-lastIdx)
			}
			lastIdx = i
		}
		pre = head
	}
	if oldIdx == lastIdx {
		return []int{-1, -1}
	}
	return []int{minDis, lastIdx - oldIdx}
}
