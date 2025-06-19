package main

import "sort"

func maxScore6(n int, edges [][]int) int64 {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	cnts := make([]int, n)
	for _, ed := range edges {
		union(ed[0], ed[1])
		cnts[ed[0]] += 1
		cnts[ed[1]] += 1
	}
	groups := map[int]int{}
	for i := 0; i < n; i++ {
		groups[find(i)] += 1
	}
	var cycles, nonCycles []int
	for i, cnt := range cnts {
		if cnt != 1 { // non cycle head/tail
			continue
		}
		groupId := find(i)
		if groups[groupId] == 0 {
			continue
		}
		nonCycles = append(nonCycles, groups[groupId])
		delete(groups, groupId)
	}
	for _, g := range groups {
		if g < 2 {
			continue
		}
		cycles = append(cycles, g)
	}
	scoreHelper := func(score int, size int, cycle bool) int {
		a := score - size + 1
		b := a + 1
		var ret int
		if size == 2 {
			return a * b
		}
		if cycle {
			ret += a * b
		}
		for a != b {
			ret += a * (a + 2)
			a += 2
			if a == score {
				ret += a * b
				break
			}
			ret += b * (b + 2)
			b += 2
			if b == score {
				ret += a * b
				break
			}
		}
		return ret
	}
	var ret int
	maxScore := n
	for _, cy := range cycles {
		ret += scoreHelper(maxScore, cy, true)
		maxScore -= cy
	}
	sort.Ints(nonCycles)
	for i := len(nonCycles) - 1; i >= 0; i-- {
		cy := nonCycles[i]
		ret += scoreHelper(maxScore, cy, false)
		maxScore -= cy
	}
	return int64(ret)
}
