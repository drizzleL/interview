package main

import (
	"math"
	"math/bits"
)

func maxGeneticDifference(parents []int, queries [][]int) []int {
	maxVal := len(parents) + 1
	for _, q := range queries {
		maxVal = max(maxVal, q[1])
	}
	bitLen := bits.Len(uint(maxVal))
	seenSize := int(math.Pow(2, float64(bitLen)+1))
	seen := make([]int, seenSize)
	var root int
	children := make([][]int, len(parents))
	for i, p := range parents {
		if p == -1 {
			root = i
			continue
		}
		children[p] = append(children[p], i)
	}
	dict := map[int][]int{}
	for i, q := range queries {
		dict[q[0]] = append(dict[q[0]], i)
	}
	mutate := func(x int, flag int) {
		var idx int
		for k := bitLen - 1; k >= 0; k-- {
			bit := 1 << k
			next := idx * 2
			if x&bit == 0 {
				next += 1
			} else {
				next += 2
			}
			seen[next] += flag
			idx = next
		}
	}
	helper := func(val int) int {
		var ret int
		var idx int
		for k := bitLen - 1; k >= 0; k-- {
			bit := 1 << k
			next := idx * 2
			switch {
			case seen[next+1] != 0 && seen[next+2] != 0:
				ret |= bit
				if val&bit == 0 {
					next += 2
				} else {
					next += 1
				}
			case seen[next+1] != 0:
				next += 1
				if val&bit != 0 {
					ret |= bit
				}
			case seen[next+2] != 0:
				next += 2
				if val&bit == 0 {
					ret |= bit
				}
			}
			idx = next
		}
		return ret
	}
	ret := make([]int, len(queries))
	var traverse func(node int)
	traverse = func(node int) {
		mutate(node, 1)
		for _, qId := range dict[node] {
			ret[qId] = helper(queries[qId][1])
		}
		for _, child := range children[node] {
			traverse(child)
		}
		mutate(node, -1)
	}
	traverse(root)
	return ret
}
