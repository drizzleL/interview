package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var q []int
	var ret []int
	in := make([]int, numCourses)
	dict := make([][]int, numCourses)
	for _, v := range prerequisites {
		in[v[0]]++
		dict[v[1]] = append(dict[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		ret = append(ret, q[0])
		for _, v := range dict[q[0]] {
			in[v]--
			if in[v] == 0 {
				q = append(q, v)
			}
		}
		q = q[1:]
	}
	if len(ret) != numCourses {
		return nil
	}
	return ret
}
