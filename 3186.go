package main

import "sort"

func maximumTotalDamage(power []int) int64 {
	sort.Ints(power)
	var powers [][2]int
	for i, p := range power {
		if i != 0 && p == power[i-1] {
			powers[len(powers)-1][1] += p
			continue
		}
		powers = append(powers, [2]int{p, p})
	}
	var q []int
	var ret int
	for _, p := range powers {
		idx := sort.Search(len(powers), func(i int) bool {
			return powers[i][0] >= p[1]-2
		})
		val := p[1]
		if idx-1 >= 0 {
			val += q[idx-1]
		}
		ret = max(ret, val)
		q = append(q, ret)
	}
	return int64(ret)
}
