package main

func numberOfAlternatingGroups2(colors []int, k int) int {
	cnt := 1
	if colors[0] != colors[len(colors)-1] {
		cnt += 1
		for i := 1; i < k-1; i++ {
			if colors[len(colors)-i] == colors[len(colors)-i-1] {
				break
			}
			cnt += 1
		}
	}
	var ret int
	if cnt >= k {
		ret = 1
	}
	for i := 1; i < len(colors); i++ {
		if colors[i] != colors[i-1] {
			cnt += 1
		} else {
			cnt = 0
		}
		if cnt >= k {
			ret += 1
		}
	}
	return ret
}
