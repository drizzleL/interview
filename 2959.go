package main

import "math"

func numberOfSets2(n int, maxDistance int, roads [][]int) int {
	distance := func(mask int, dist [][]int) int {
		for k := 0; k < n; k++ {
			if (1<<k)&mask == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if (1<<i)&mask == 0 || i == k {
					continue
				}
				for j := 0; j < n; j++ {
					if (1<<j)&mask == 0 || j == k || j == i {
						continue
					}
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
		var ret int
		for i := 0; i < n; i++ {
			if (1<<i)&mask == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (1<<j)&mask == 0 {
					continue
				}
				ret = max(ret, dist[i][j])
			}
		}
		return ret
	}
	var ret int
	for mask := 0; mask < 1<<n; mask++ {
		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := range dist[i] {
				dist[i][j] = math.MaxInt32
			}
			dist[i][i] = 0
		}
		for _, r := range roads {
			dist[r[0]][r[1]] = min(dist[r[0]][r[1]], r[2])
			dist[r[1]][r[0]] = min(dist[r[1]][r[0]], r[2])
		}
		if distance(mask, dist) <= maxDistance {
			ret += 1
		}
	}
	return ret
}
