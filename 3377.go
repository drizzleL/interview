package main

import (
	"container/heap"
	"math"
)

func minOperations5(n int, m int) int {
	var digits int
	for k := max(n, m); k != 0; digits, k = digits+1, k/10 {
	}
	primes := make([]bool, int(math.Pow10(digits)))
	primes[0] = true
	primes[1] = true
	for i := 2; i < len(primes); i++ {
		if primes[i] {
			continue
		}
		for j := 2; j*i < len(primes); j++ {
			primes[i*j] = true
		}
	}
	if !primes[m] || !primes[n] {
		return -1
	}
	q := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[1] < b.([2]int)[1]
		},
	}
	heap.Push(&q, [2]int{n, n})
	for q.Len() != 0 {
		top := heap.Pop(&q).([2]int)
		if top[0] == m {
			return top[1]
		}
		if !primes[top[0]] {
			continue
		}
		for base, bit := 1, 1; base < len(primes); base, bit = base*10, bit+1 {
			num := top[0] / base % 10
			if num > 1 || num == 1 && bit != digits {
				heap.Push(&q, [2]int{top[0] - base, top[1] + top[0] - base})
			}
			if num < 9 {
				heap.Push(&q, [2]int{top[0] + base, top[1] + top[0] + base})
			}
		}
		primes[top[0]] = false
	}
	return -1
}
