package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	prers := make([][]int, numCourses)
	for _, prer := range prerequisites {
		prers[prer[0]] = append(prers[prer[0]], prer[1])
	}
	children := make([][]bool, numCourses)
	for i := range children {
		children[i] = make([]bool, numCourses)
	}
	cache := make([]bool, numCourses)
	ans := make([][]int, numCourses)
	var getChildren func(i int) []int
	getChildren = func(i int) (ret []int) {
		if cache[i] {
			return ans[i]
		}
		defer func() {
			cache[i] = true
			ans[i] = ret
		}()
		seen := make([]bool, numCourses)
		for _, directChild := range prers[i] {
			seen[directChild] = true
			for _, c := range getChildren(directChild) {
				seen[c] = true
			}
		}
		for child, ok := range seen {
			if ok {
				seen[child] = true
				ret = append(ret, child)
			}
		}
		return
	}
	for i := 0; i < numCourses; i++ {
		for _, child := range getChildren(i) {
			children[i][child] = true
		}
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		ret[i] = children[q[0]][q[1]]
	}
	return ret
}
