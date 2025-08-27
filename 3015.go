package main

import "fmt"

func countOfPairs2(n int, x int, y int) []int {
	dict := make([][]int, n+1)
	ret := make([]int, n)
	for i := 1; i < n; i++ {
		dict[i] = append(dict[i], i+1)
		dict[i+1] = append(dict[i+1], i)
	}
	if x > y {
		x, y = y, x
	}
	if x != y && x+1 != y {
		dict[x] = append(dict[x], y)
		dict[y] = append(dict[y], x)
	}
	for i := 1; i <= n; i++ {
		seen := make([]bool, n+1)
		nodes := []int{i}
		for step := 0; len(nodes) != 0; step++ {
			var next []int
			for _, node := range nodes {
				seen[node] = true
				for _, child := range dict[node] {
					if seen[child] {
						continue
					}
					fmt.Println(node, child, step)
					ret[step] += 1
					seen[child] = true
					next = append(next, child)
				}
			}
			nodes = next
		}
	}
	return ret
}
