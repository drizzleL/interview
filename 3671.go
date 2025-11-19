package main

import "sort"

func totalBeauty(nums []int) int {
	var maxVal int
	dict := map[int][]int{}
	for i, num := range nums {
		maxVal = max(maxVal, num)
		dict[num] = append(dict[num], i)
	}
	dp := make([]int, maxVal+1)
	for i := 1; i < len(dp); i++ {
		var idxs []int
		idxDict := map[int]int{}
		for d := i; d <= maxVal; d += i {
			if len(dict[d]) == 0 {
				continue
			}
			idxDict[d] = len(idxDict)
			idxs = append(idxs, dict[d]...)
		}
		sort.Ints(idxs)
		f := NewFenwick(len(idxDict))
		for _, idx := range idxs {
			num := nums[idx]
			pos := idxDict[num]
			presum := f.sum(pos - 1)
			f.add(pos, presum+1)
		}
		dp[i] += f.sum(len(idxDict) - 1)
		dp[i] %= 1e9 + 7
	}
	for d := maxVal; d > 0; d-- {
		for e := 2 * d; e <= maxVal; e += d {
			dp[d] = (dp[d] - dp[e] + 1e9 + 7) % (1e9 + 7)
		}
	}
	var ret int
	for i := 1; i < len(dp); i++ {
		ret += i * dp[i]
		ret %= 1e9 + 7
	}
	return ret
}

type Fenwick struct {
	a []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		a: make([]int, n+1),
	}
}

func (f *Fenwick) sum(i int) int {
	i++
	s := 0
	for i > 0 {
		s += f.a[i]
		i -= i & -i
	}
	return s
}

func (f *Fenwick) add(i, x int) {
	i++
	for i < len(f.a) {
		f.a[i] += x
		i += i & -i
	}
}
