package main

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	var ret int
	sort.Ints(coins)
	var idx int
	for k := 0; k < target; {
		if idx < len(coins) && coins[idx] <= k+1 {
			k += coins[idx]
			idx += 1
		} else {
			k += k + 1
			ret += 1
		}
	}
	return ret
}
