package main

func maxProfit10(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	dict := make([][]int, n)
	dict2 := make([][]int, n)
	out := make([]int, n)
	for _, h := range hierarchy {
		out[h[0]-1] += 1
		dict[h[0]-1] = append(dict[h[0]-1], h[1]-1)
		dict2[h[1]-1] = append(dict2[h[1]-1], h[0]-1)
	}
	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, 2)
		for j := range cache[i] {
			cache[i][j] = make([]int, budget+1)
		}
	}
	var nodes []int
	for i := 0; i < n; i++ {
		if out[i] == 0 {
			nodes = append(nodes, i)
		}
	}
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		for _, next := range dict2[node] {
			out[next] -= 1
			if out[next] == 0 {
				nodes = append(nodes, next)
			}
		}
		for j := 0; j <= 1; j++ {
			base := make([]int, budget+1)
			for _, child := range dict[node] {
				next := make([]int, budget+1)
				for b1 := 0; b1 <= budget; b1++ {
					for b2 := 0; b1+b2 <= budget; b2++ {
						next[b1+b2] = max(next[b1+b2], base[b1]+cache[child][0][b2])
					}
				}
				base = next
			}
			for i := range base {
				cache[node][j][i] = max(cache[node][j][i], base[i])
			}

			price := present[node]
			if j == 1 {
				price = present[node] / 2
			}
			profit := future[node] - price
			if price <= budget {
				base := make([]int, budget+1)
				for _, child := range dict[node] {
					next := make([]int, budget+1)
					for b1 := 0; b1 <= budget; b1++ {
						for b2 := 0; b1+b2 <= budget; b2++ {
							next[b1+b2] = max(next[b1+b2], base[b1]+cache[child][1][b2])
						}
					}
					base = next
				}
				for i := price; i <= budget; i++ {
					cache[node][j][i] = max(cache[node][j][i], base[i-price]+profit)
				}
			}
		}
	}
	return cache[0][0][budget]
}
