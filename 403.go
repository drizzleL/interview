package main

import (
	"log"
	"sort"
)

func canCross(stones []int) bool {
	cache := map[[2]int]bool{}
	try := func(i int, step int) int {
		if step <= 0 {
			return -1
		}
		next := stones[i] + step
		idx := sort.SearchInts(stones[i+1:], next)
		if idx+i+1 == len(stones) {
			return -1
		}
		if stones[idx+i+1] != next {
			return -1
		}
		return idx + i + 1
	}
	var helper func(i int, step int) bool
	helper2 := func(i, step int) bool {
		next := try(i, step)
		if next == -1 {
			return false
		}
		return helper(next, step)
	}
	helper = func(i int, step int) (ret bool) {
		if i == len(stones)-1 {
			log.Println(i, step)
			return true
		}
		if c, ok := cache[[2]int{i, step}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, step}] = ret
		}()
		return helper2(i, step-1) || helper2(i, step) || helper2(i, step+1)
	}
	return helper(0, 0)
}
