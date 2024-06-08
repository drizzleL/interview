package main

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	dict := map[[2]int][][2]int{}
	for i := 0; i < len(ring); i++ {
		dict[[2]int{i, int(ring[i] - 'a')}] = append(dict[[2]int{i, int(ring[i] - 'a')}], [2]int{i, 0})
		l, r := i, i
		for v := 1; ; v++ {
			l -= 1
			r += 1
			if l < 0 {
				l += len(ring)
			}
			if r >= len(ring) {
				r -= len(ring)
			}
			if ring[l] != ring[i] {
				dict[[2]int{i, int(ring[l] - 'a')}] = append(dict[[2]int{i, int(ring[l] - 'a')}], [2]int{l, v})
			}
			if l != r && ring[r] != ring[i] {
				dict[[2]int{i, int(ring[r] - 'a')}] = append(dict[[2]int{i, int(ring[r] - 'a')}], [2]int{r, v})
			}
			if l == r || l+1 == r {
				break
			}
		}
	}
	cache := map[[2]int]int{}
	var helper func(keyIdx int, pos int) int
	helper = func(keyIdx int, pos int) (ret int) {
		if keyIdx == len(key) {
			return 0
		}
		if c, ok := cache[[2]int{keyIdx, pos}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{keyIdx, pos}] = ret
		}()
		nextCh := int(key[keyIdx] - 'a')
		ret = math.MaxInt32
		for _, v := range dict[[2]int{pos, nextCh}] {
			nextPos, takes := v[0], v[1]
			ret = min(ret, helper(keyIdx+1, nextPos)+takes)
		}
		return
	}
	return helper(0, 0) + len(key)
}
