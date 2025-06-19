package main

import (
	"sort"
)

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	for i := 1; i < len(fruits); i++ {
		fruits[i][1] += fruits[i-1][1]
	}
	var ret int
	for i := 0; i < len(fruits) && fruits[i][0] <= startPos; i++ {
		pos := fruits[i][0]
		if startPos-pos > k { // can't reach
			continue
		}
		var pre int
		if i != 0 {
			pre = fruits[i-1][1]
		}
		var midPos int
		if (startPos-pos)*2 >= k { // can't go back
			midPos = sort.Search(len(fruits), func(i int) bool {
				return fruits[i][0] >= startPos+1
			}) - 1
		} else {
			midPos = sort.Search(len(fruits), func(i int) bool {
				return fruits[i][0] >= startPos+(k-(startPos-pos)*2)+1
			}) - 1
		}
		ret = max(ret, fruits[midPos][1]-pre)
	}
	for i := len(fruits) - 1; i >= 0 && fruits[i][0] >= startPos; i-- {
		pos := fruits[i][0]
		if pos-startPos > k {
			continue
		}
		var pre int
		var midPos int
		if (pos-startPos)*2 >= k {
			midPos = sort.Search(len(fruits), func(i int) bool {
				return fruits[i][0] >= startPos
			}) - 1
		} else {
			midPos = sort.Search(len(fruits), func(i int) bool {
				return fruits[i][0] >= startPos-(k-(pos-startPos)*2)
			}) - 1
		}
		if midPos >= 0 {
			pre = fruits[midPos][1]
		}
		ret = max(ret, fruits[i][1]-pre)
	}
	return ret
}
