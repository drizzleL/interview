package main

import (
	"fmt"
	"math/bits"
)

func kthSmallest2(par []int, vals []int, queries [][]int) []int {
	ret := make([]int, len(queries))
	queryDict := map[int][]int{}
	var maxBits int
	for _, v := range vals {
		maxBits = max(maxBits, bits.Len(uint(v)))
	}
	for i, q := range queries {
		queryDict[q[0]] = append(queryDict[q[0]], i)
	}
	childDict := make([][]int, len(par))
	for i, p := range par {
		if p == -1 {
			continue
		}
		childDict[p] = append(childDict[p], i)
	}
	var dfs func(x int, p int) *TrieNode
	dfs = func(x int, p int) *TrieNode {
		p ^= vals[x]
		t := &TrieNode{}
		addToTrie(t, p, maxBits)
		for _, child := range childDict[x] {
			t = mergeTrie(t, dfs(child, p))
		}
		if _, ok := queryDict[x]; ok {
			for _, qId := range queryDict[x] {
				k := queries[qId][1]
				ret[qId] = findKth(t, k, maxBits)
			}
		}
		return t
	}
	dfs(0, 0)
	return ret
}

type TrieNode struct {
	children [2]*TrieNode
	cnt      int
}

func printTrie(b *TrieNode, maxBits int, v int) {
	if maxBits == 0 {
		fmt.Println(v)
		return
	}
	if b.children[0] != nil {
		printTrie(b.children[0], maxBits-1, v)
	}
	if b.children[1] != nil {
		printTrie(b.children[1], maxBits-1, v+1<<(maxBits-1))
	}
}

func addToTrie(b *TrieNode, v int, maxBits int) {
	b.cnt += 1
	for i := maxBits - 1; i >= 0; i-- {
		bits := (v >> i) & 1
		if b.children[bits] == nil {
			b.children[bits] = &TrieNode{}
		}
		b = b.children[bits]
		b.cnt += 1
	}
}

func mergeTrie(base, b *TrieNode) *TrieNode {
	if base == nil {
		return b
	}
	if b == nil {
		return base
	}
	if base.cnt < b.cnt {
		base, b = b, base
	}
	base.children[0] = mergeTrie(base.children[0], b.children[0])
	base.children[1] = mergeTrie(base.children[1], b.children[1])
	base.cnt = 0
	if base.children[0] != nil {
		base.cnt += base.children[0].cnt
	}
	if base.children[1] != nil {
		base.cnt += base.children[1].cnt
	}
	base.cnt = max(base.cnt, 1)
	return base
}

func findKth(b *TrieNode, k int, maxBits int) int {
	if b.cnt < k {
		return -1
	}
	var ret int
	for i := maxBits - 1; i >= 0; i-- {
		if b.children[0] != nil && b.children[0].cnt >= k {
			b = b.children[0]
			continue
		}
		if b.children[0] != nil {
			k -= b.children[0].cnt
		}
		ret += 1 << i
		b = b.children[1]
	}
	return ret
}
