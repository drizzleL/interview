package main

func score(cards []string, x byte) int {
	size := int('j'-'a') + 1
	a, b := make([]int, size), make([]int, size)
	var both int
	for _, c := range cards {
		if c[0] != x && c[1] != x {
			continue
		}
		if c[0] == x && c[1] == x {
			both += 1
			continue
		}
		if c[0] == x {
			a[c[1]-'a'] += 1
		} else {
			b[c[0]-'a'] += 1
		}
	}
	helper := func(dict []int, both int) (cnt int, ret int, both2 int) {
		var maxVal int
		for _, v := range dict {
			cnt += v
			maxVal = max(maxVal, v)
		}
		if cnt >= maxVal*2 { // remains
			return cnt, 0, both
		}
		shouldConsume := maxVal - cnt/2
		consume := min(both, shouldConsume)
		cnt -= shouldConsume
		both -= consume
		return cnt, consume, both
	}
	var ret int
	var cnt1, cnt2, tmp int
	cnt1, tmp, both = helper(a, both)
	ret += tmp
	cnt2, tmp, both = helper(b, both)
	ret += tmp
	helper2 := func(cnt int, both int) (ret int, both2 int) {
		if both >= cnt {
			return cnt, both - cnt
		}
		ret = (cnt + both) / 2
		both -= cnt + both - 2*ret
		return ret, both
	}
	tmp, both = helper2(cnt1, both)
	ret += tmp
	tmp, both = helper2(cnt2, both)
	ret += tmp
	return ret
}
