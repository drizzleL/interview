package main

import "sort"

func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}
	helper := func(x []int) int {
		var ret int
		sorted := make([]int, len(x))
		copy(sorted, x)
		sort.Ints(sorted)
		dict := map[int]int{}
		for i, v := range sorted {
			dict[v] = i
		}
		for i := 0; i < len(x); {
			if x[i] == sorted[i] {
				i++
				continue
			}
			j := dict[x[i]]
			x[i], x[j] = x[j], x[i]
			ret += 1
		}
		return ret
	}
	nodes := []*TreeNode{root}
	var ret int
	for len(nodes) != 0 {
		var next []*TreeNode
		var vals []int
		for _, n := range nodes {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
			vals = append(vals, n.Val)
		}
		ret += helper(vals)
	}
	return ret
}
