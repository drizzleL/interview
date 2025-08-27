package main

func countTrapezoids(points [][]int) int {
	dict := map[int]int{}
	for _, p := range points {
		dict[p[1]] += 1
	}
	var nodes []int
	for _, v := range dict {
		if v == 1 {
			continue
		}
		nodes = append(nodes, v*(v-1)/2)
	}
	var ret int
	var sum int
	for _, node := range nodes {
		ret += sum * node
		ret %= 1e9 + 7
		sum += node
	}
	return ret
}
