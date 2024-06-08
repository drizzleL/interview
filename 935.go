package main

func knightDialer(n int) int {
	dict := map[int][]int{
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {3, 9, 0},
		5: {},
		6: {1, 0, 7},
		7: {2, 6},
		8: {1, 3},
		9: {2, 4},
		0: {4, 6},
	}
	var now, next [10]int
	for i := range now {
		now[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 10; j++ {
			next[j] = 0
			for _, v := range dict[j] {
				next[j] += now[v]
			}
			next[j] %= 1e9 + 7
		}
		now, next = next, now
	}
	var ret int
	for i := range now {
		ret += now[i]
		ret %= 1e9 + 7
	}
	return ret
}
