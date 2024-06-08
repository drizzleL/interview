package main

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	maxBuild := func(machine int) int {
		comp := composition[machine]
		var l, r int
		for i, c := range comp {
			if c == 0 {
				continue
			}
			r = max(r, (budget/cost[i]+stock[i])/c)
		}
		for l < r {
			mid := (l + r) / 2
			if mid == l {
				mid += 1
			}
			var needMoney int
			for i, c := range comp {
				if stock[i] >= mid*c {
					continue
				}
				needMoney += cost[i] * (mid*c - stock[i])
			}
			if budget >= needMoney {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return l
	}
	var ret int
	for i := range composition {
		ret = max(ret, maxBuild(i))
	}
	return ret
}
