package main

func equationsPossible(equations []string) bool {
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
	for _, eq := range equations {
		if eq[1:3] != "==" {
			continue
		}
		a, b := eq[0], eq[3]
		union(int(a-'a'), int(b-'a'))
	}
	for _, eq := range equations {
		if eq[1:3] != "!=" {
			continue
		}
		a, b := eq[0], eq[3]
		if find(int(a-'a')) == find(int(b-'a')) {
			return false
		}
	}
	return true
}
