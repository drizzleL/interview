package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	getPath := func(x int) []int {
		var ret []int
		seen := make([]bool, len(edges))
		for x != -1 && !seen[x] {
			ret = append(ret, x)
			seen[x] = true
			x = edges[x]
		}
		return ret
	}
	p1, p2 := getPath(node1), getPath(node2)
	s1, s2 := make([]bool, len(edges)), make([]bool, len(edges))
	ret := -1
	for ret == -1 && (len(p1) != 0 || len(p2) != 0) {
		if len(p1) != 0 {
			node := p1[0]
			if s2[node] {
				if ret == -1 || node < ret {
					ret = node
				}
			}
			s1[node] = true
			p1 = p1[1:]
		}
		if len(p2) != 0 {
			node := p2[0]
			if s1[node] {
				if ret == -1 || node < ret {
					ret = node
				}
			}
			s2[node] = true
			p2 = p2[1:]
		}
	}
	return ret
}
