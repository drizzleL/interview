package main

func distinctPoints(s string, k int) int {
	helper := func(c byte) int {
		switch c {
		case 'U':
			return 0
		case 'D':
			return 1
		case 'L':
			return 2
		case 'R':
			return 3
		}
		return 0
	}
	var total, curr [4]int
	for _, c := range s {
		total[helper(byte(c))] += 1
	}
	for i := 0; i < k; i++ {
		curr[helper(s[i])] += 1
	}
	dict := map[[2]int]bool{}
	add := func() {
		up, down, left, right := total[0]-curr[0], total[1]-curr[1], total[2]-curr[2], total[3]-curr[3]
		key := [2]int{up - down, left - right}
		dict[key] = true
	}
	add()
	for i := k; i < len(s); i++ {
		curr[helper(s[i])] += 1
		curr[helper(s[i-k])] -= 1
		add()
	}
	return len(dict)
}
