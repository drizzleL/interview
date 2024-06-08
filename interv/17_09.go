package main

import (
	"math"
)

func getKthMagicNumber(k int) int {
	if k <= 1 {
		return k
	}
	data := []int{1}
	a, b, c := 0, 0, 0
	for i := 1; i < k; i++ {
		val := math.MaxInt32
		val = min(val, data[a]*3)
		val = min(val, data[b]*5)
		val = min(val, data[c]*7)
		if val == data[a]*3 {
			a++
		}
		if val == data[b]*5 {
			b++
		}
		if val == data[c]*7 {
			c++
		}
		data = append(data, val)
	}
	return data[len(data)-1]
}
