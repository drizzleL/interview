package main

import (
	"sort"
)

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	var sumList []int
	for len(nodes) > 0 {
		var next []*TreeNode
		var sum int
		for _, n := range nodes {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		sumList = append(sumList, sum)
		nodes = next
	}
	if len(sumList) < k {
		return -1
	}
	sort.Ints(sumList)
	return int64(sumList[len(sumList)-k])
}
