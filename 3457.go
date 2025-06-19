package main

import (
	"sort"
)

func maxWeight(pizzas []int) int64 {
	sort.Ints(pizzas)
	days := len(pizzas) / 4
	oddDays := days/2 + days%2
	evenDays := days - oddDays
	var ret int
	for i := 0; i < oddDays; i++ {
		ret += pizzas[len(pizzas)-1-i]
	}
	for i := 0; i < evenDays; i++ {
		ret += pizzas[len(pizzas)-oddDays-2-i*2]
	}
	return int64(ret)
}
