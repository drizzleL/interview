package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	dict := map[int][]int{}
	in := make([]int, numCourses)
	for _, p := range prerequisites {
		u, v := p[0], p[1]
		in[v] += 1
		dict[u] = append(dict[u], v)
	}
	parents := map[int]map[int]bool{}
	for i := 0; i < numCourses; i++ {
		parents[i] = map[int]bool{
			i: true,
		}
	}
	var nodes []int
	for i, v := range in {
		if v == 0 {
			nodes = append(nodes, i)
		}
	}
	for len(nodes) != 0 {
		var next []int
		for _, n := range nodes {
			for _, child := range dict[n] {
				for p := range parents[n] {
					parents[child][p] = true
				}
				in[child] -= 1
				if in[child] == 0 {
					next = append(next, child)
				}
			}
		}
		nodes = next
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		ret[i] = parents[v][u]
	}
	return ret
}
