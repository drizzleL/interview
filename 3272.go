package main

import (
	"math"
	"strconv"
)

func countGoodIntegers(n int, k int) int64 {
	var ret int
	ways := map[[10]int]int{}
	getStr := func(num int) string {
		str := strconv.Itoa(num)
		b := make([]byte, n)
		for i := range b {
			b[i] = '0'
		}
		for i, j, k := 0, len(b)-1, len(str)-1; k >= 0; i, j, k = i+1, j-1, k-1 {
			b[i] = str[k]
			b[j] = str[k]
		}
		return string(b)
	}
	divisible := func(str string) bool {
		v, _ := strconv.Atoi(str)
		return v%k == 0
	}
	getKey := func(str string) [10]int {
		key := [10]int{}
		for _, c := range str {
			key[c-'0'] += 1
		}
		return key
	}
	getWays := func(key [10]int) int {
		ret := 1
		size := n
		ret *= combination(size-1, key[0])
		size -= key[0]
		for i := 1; i <= 9; i++ {
			ret *= combination(size, key[i])
			size -= key[i]
		}
		return ret
	}
	maxNum := int(math.Pow10((n + 1) / 2))
	for num := 1; num < maxNum; num++ {
		if num%10 == 0 {
			continue
		}
		str := getStr(num)
		key := getKey(str)
		if ways[key] != 0 { // checked before
			continue
		}
		if !divisible(str) {
			continue
		}
		way := getWays(key)
		ret += way
		ways[key] = way
	}
	return int64(ret)
}
