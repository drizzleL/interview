package main

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			parent[a] = find(parent[a])
		}
		return parent[a]
	}
	union := func(a, b int) {
		if a == b {
			return
		}
		if find(a) > find(b) {
			a, b = b, a
		}
		parent[find(b)] = find(a)
	}
	for i := range s1 {
		c1, c2 := s1[i], s2[i]
		union(int(c1-'a'), int(c2-'a'))
	}
	ret := make([]byte, 0, len(baseStr))
	for _, c := range baseStr {
		ret = append(ret, byte('a'+find(int(c-'a'))))
	}
	return string(ret)
}
