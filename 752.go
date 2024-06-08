package main

func openLock(deadends []string, target string) int {
	dead := map[[4]int]bool{}
	for _, d := range deadends {
		var tmp [4]int
		for i := range d {
			tmp[i] = int(d[i] - '0')
		}
		dead[tmp] = true
	}
	var t [4]int
	for i := range target {
		t[i] = int(target[i] - '0')
	}
	var curr [4]int
	if curr == t {
		return 0
	}
	if dead[curr] {
		return -1
	}
	seen := map[[4]int]bool{}
	nodes := [][4]int{curr}
	try := func(x [4]int, i int) [][4]int {
		var next [][4]int
		if x[i] == 0 {
			x[i] = 9
			next = append(next, x)
			x[i] = 1
			next = append(next, x)
		} else if x[i] == 9 {
			x[i] = 8
			next = append(next, x)
			x[i] = 0
			next = append(next, x)
		} else {
			x[i] += 1
			next = append(next, x)
			x[i] -= 2
			next = append(next, x)
		}
		return next
	}
	for round := 1; len(nodes) != 0; round++ {
		var next [][4]int
		for _, n := range nodes {
			for i := range n {
				for _, can := range try(n, i) {
					if can == t {
						return round
					}
					if !dead[can] && !seen[can] {
						next = append(next, can)
						seen[can] = true
					}
				}
			}
		}
		nodes = next
	}
	return -1
}
