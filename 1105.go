package main

import (
	"math"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	cache := map[int]int{}
	var helper func(i int) int
	helper = func(i int) (ret int) {
		if c, ok := cache[i]; ok {
			return c
		}
		defer func() {
			cache[i] = ret
		}()
		if i == len(books) {
			return 0
		}
		ret = math.MaxInt32
		maxH := books[i][1]
		width := books[i][0]
		j := i + 1
		for {
			for j < len(books) && books[j][1] <= maxH && width+books[j][0] <= shelfWidth {
				width += books[j][0]
				j++
			}
			if j == len(books) || width+books[j][0] > shelfWidth {
				ret = min(ret, maxH+helper(j))
				break
			}
			ret = min(ret, maxH+helper(j))
			maxH = books[j][1]
			width += books[j][0]
			j += 1
		}
		return
	}
	return helper(0)
}
