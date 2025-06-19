package main

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
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
	for i := range s1 {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	b := []byte(baseStr)
	for i := range b {
		b[i] = byte('a' + find(int(b[i]-'a')))
	}
	return string(b)
}
