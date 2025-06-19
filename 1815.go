package main

import (
	"strconv"
	"strings"
)

func maxHappyGroups(batchSize int, groups []int) int {
	dict := make([]int, batchSize)
	var ret int
	for _, g := range groups {
		m := g % batchSize
		if m == 0 {
			ret += 1
			continue
		}
		if dict[batchSize-m] > 0 {
			dict[batchSize-m]--
			ret += 1
			continue
		}
		dict[g%batchSize]++
	}
	getKey := func(cnt []int, remainder int) string {
		var strs []string
		for i := 1; i < batchSize; i++ {
			strs = append(strs, strconv.Itoa(cnt[i]))
		}
		strs = append(strs, strconv.Itoa(remainder))
		return strings.Join(strs, ",")
	}
	cache := map[string]int{}
	var helper func(cnt []int, remainder int) int
	helper = func(cnt []int, remainder int) (ret int) {
		key := getKey(cnt, remainder)
		if c, ok := cache[key]; ok {
			return c
		}
		defer func() {
			cache[key] = ret
		}()
		if remainder > 0 && cnt[batchSize-remainder] > 0 {
			cnt[batchSize-remainder]--
			ret = helper(cnt, 0)
			cnt[batchSize-remainder]++
			return ret
		}
		for i := 1; i < batchSize; i++ {
			if cnt[i] == 0 {
				continue
			}
			cnt[i]--
			tmp := helper(cnt, (remainder+i)%batchSize)
			if remainder == 0 {
				tmp += 1
			}
			ret = max(ret, tmp)
			cnt[i]++
		}
		return
	}
	return ret + helper(dict, 0)
}
