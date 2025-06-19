package main

import (
	"sort"
)

func maxProfit3(inventory []int, orders int) int {
	sort.Ints(inventory)
	var ret int
	curr := inventory[len(inventory)-1]
	for i := len(inventory) - 1; i >= 0; {
		for i >= 0 && inventory[i] == curr {
			i -= 1
		}
		var pre int
		if i >= 0 {
			pre = inventory[i]
		}
		size := len(inventory) - i - 1
		maxSize := (curr - pre) * size
		if maxSize >= orders { // can take all orders
			levels := orders / size
			ret += (curr + curr - levels + 1) * levels / 2 * size
			ret %= 1e9 + 7
			left := orders - levels*size
			ret += left * (curr - levels)
			ret %= 1e9 + 7
			break
		}
		orders -= maxSize
		ret += (pre + 1 + curr) * (curr - pre) / 2 * size
		ret %= 1e9 + 7
		curr = inventory[i]
	}
	return ret
}
