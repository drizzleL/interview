package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var helper func(i int, curr []int)
	var ret [][]int
	helper = func(i int, curr []int) {
		curr = append(curr, i)
		if i == len(graph)-1 {
			ret = append(ret, append([]int{}, curr...))
			return
		}
		for _, e := range graph[i] {
			helper(e, curr)
		}
	}
	helper(0, nil)
	return ret
}
