package main

import (
	"sort"
)

func maximumScore(nums []int, k int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]int, maxVal+1)
	for i := 2; i < len(sieve); i++ {
		if sieve[i] != 0 {
			continue
		}
		sieve[i] = 1
		for j := 2; j*i < len(sieve); j++ {
			sieve[i*j] += 1
		}
	}

	type Ele struct {
		Idx int
		Num int
	}
	var eles []*Ele
	for i, num := range nums {
		eles = append(eles, &Ele{
			Idx: i,
			Num: num,
		})
	}
	sort.Slice(eles, func(i, j int) bool {
		if eles[i].Num == eles[j].Num {
			return eles[i].Idx < eles[j].Idx
		}
		return eles[i].Num > eles[j].Num
	})
	ret := 1
	helper := func(x int, m, p int) int {
		for p != 0 {
			if p&1 == 1 {
				x *= m
				x %= 1e9 + 7
			}
			m *= m
			m %= 1e9 + 7
			p >>= 1
		}
		return x
	}
	befores, afters := make([]int, len(nums)), make([]int, len(nums))
	var stack []int
	for i, num := range nums {
		for len(stack) != 0 && sieve[nums[stack[len(stack)-1]]] < sieve[num] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			befores[i] = -1
		} else {
			befores[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for j := len(nums) - 1; j >= 0; j-- {
		num := nums[j]
		for len(stack) != 0 && sieve[nums[stack[len(stack)-1]]] <= sieve[num] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			afters[j] = len(nums)
		} else {
			afters[j] = stack[len(stack)-1]
		}
		stack = append(stack, j)
	}
	for _, ele := range eles {
		before, after := befores[ele.Idx], afters[ele.Idx]
		poss := (ele.Idx - before) * (after - ele.Idx)
		if poss >= k {
			ret = helper(ret, ele.Num, k)
			break
		}
		ret = helper(ret, ele.Num, poss)
		k -= poss
	}
	return ret
}
