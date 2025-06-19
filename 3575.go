package main

func goodSubtreeSum(vals []int, par []int) int {
	dict := make([][1024]int, len(vals))
	getKey := func(v int) int {
		var ret int
		for v != 0 {
			d := v % 10
			if (1<<d)&ret != 0 {
				return -1
			}
			ret |= 1 << d
			v /= 10
		}
		return ret
	}
	for i, v := range vals {
		k := getKey(v)
		if k == -1 {
			continue
		}
		dict[i][k] = max(dict[i][k], v)
		for j := i; j != -1; j = par[j] {
			for k2, v2 := range dict[j] {
				if k2&k != 0 {
					continue
				}
				dict[j][k2|k] = max(dict[j][k2|k], v2+v)
				dict[j][k2|k] %= 1e9 + 7
			}
		}
	}
	var ret int
	for i := range vals {
		var tmp int
		for _, v := range dict[i] {
			tmp = max(tmp, v)
		}
		ret += tmp
		ret %= 1e9 + 7
	}
	return ret
}
