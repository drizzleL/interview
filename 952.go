package main

func largestComponentSize(nums []int) int {
	primeMask, numMask := 1<<0, 1<<1
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	parent := make([]int, maxVal+1)
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
	sieve := make([]int, maxVal+1)
	for _, num := range nums {
		sieve[num] |= numMask
	}
	for i := 2; i <= maxVal; i++ {
		if sieve[i]&primeMask != 0 { // skip non prime
			continue
		}
		for j := i * 2; j <= maxVal; j += i {
			if sieve[j]&numMask != 0 { //
				union(i, j)
			}
			sieve[j] |= primeMask
		}
	}
	var ret int
	group := map[int]int{}
	for _, num := range nums {
		p := find(num)
		group[p] += 1
		ret = max(ret, group[p])
	}
	return ret
}
