package main

import (
	"container/heap"
	"math"
	"math/bits"
)

func minTime2(n int, k int, m int, time []int, mul []float64) float64 {
	type state struct {
		mask  int
		stage int
		p     int
		t     float64
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(state).t < b.(state).t
		},
	}
	seen := make([][][]bool, 1<<n)
	for i := range seen {
		seen[i] = make([][]bool, m)
		for j := range seen[i] {
			seen[i][j] = make([]bool, 2)
		}
	}
	heap.Push(h, state{0, 0, 0, 0})
	endMask := 1<<n - 1
	for h.Len() > 0 {
		top := heap.Pop(h).(state)
		if top.mask == endMask {
			return top.t
		}
		if seen[top.mask][top.stage][top.p] {
			continue
		}
		seen[top.mask][top.stage][top.p] = true
		if top.p == 1 { // pick one return
			for i := 0; i < n; i++ {
				if top.mask&(1<<i) == 0 { // cant chosen
					continue
				}
				t2 := float64(time[i]) * mul[top.stage]
				nextStage := (top.stage + int(t2)) % m
				state2 := state{top.mask ^ (1 << i), nextStage, 0, top.t + t2}
				if seen[state2.mask][state2.stage][state2.p] {
					continue
				}
				heap.Push(h, state2)
			}
			continue
		}
		var cand []int
		for i := 0; i < n; i++ {
			if top.mask&(1<<i) != 0 { // cant chosen
				continue
			}
			cand = append(cand, i)
		}
		f := generateSubsets(cand, min(k, len(cand)))
		for {
			g := f()
			if g == nil {
				break
			}
			var s int
			mask2 := top.mask
			for _, v := range g {
				mask2 |= 1 << v
				s = max(s, time[v])
			}
			t2 := float64(s) * mul[top.stage]
			nextStage := (top.stage + int(t2)) % m
			state2 := state{mask2, nextStage, 1, top.t + t2}
			if seen[state2.mask][state2.stage][state2.p] {
				continue
			}
			heap.Push(h, state2)
		}
	}
	return -1
}

func minTime3(n int, k int, m int, time []int, mul []float64) float64 {
	// 如果只能采集一个物品但有多个物品，问题无解
	if k == 1 && n > 1 {
		return -1
	}

	// 预计算每个物品集合的最大时间
	maxTime := make([]int, 1<<n)
	for i := range maxTime {
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				if time[j] > maxTime[i] {
					maxTime[i] = time[j]
				}
			}
		}
	}

	// 初始化DP数组
	dp := make([][][][]float64, 1<<n)
	for i := range dp {
		dp[i] = make([][][]float64, m)
		for j := range dp[i] {
			dp[i][j] = make([][]float64, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]float64, 4)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	// 递归函数
	var dfs func(uint, int, int, int) float64
	dfs = func(mask uint, st, across, singles int) float64 {
		if mask == 0 {
			return 0
		}
		if singles > 3 {
			return math.MaxFloat64
		}
		if dp[mask][st][across][singles] == -1 {
			res := math.MaxFloat64
			if across == 0 { // 去程：采集物品
				// 枚举所有非空子集
				for i := uint(1); i <= mask; i++ {
					if i&mask == i && bits.OnesCount(i) <= k {
						took := float64(maxTime[i]) * mul[st]
						nextSt := (st + int(math.Floor(took))) % m
						newSingles := singles
						if bits.OnesCount(i) == 1 {
							newSingles++
						}
						candidate := took + dfs(mask-i, nextSt, 1-across, newSingles)
						if candidate < res {
							res = candidate
						}
					}
				}
			} else { // 返程：放回物品
				for i := 0; i < n; i++ {
					if mask&(1<<i) == 0 {
						took := float64(time[i]) * mul[st]
						nextSt := (st + int(math.Floor(took))) % m
						candidate := took + dfs(mask|(1<<i), nextSt, 1-across, singles)
						if candidate < res {
							res = candidate
						}
					}
				}
			}
			dp[mask][st][across][singles] = res
		}
		return dp[mask][st][across][singles]
	}

	return dfs(uint(1<<n-1), 0, 0, 0)
}

func generateSubsets(cand []int, k int) func() []int {
	n := len(cand)
	stack := []int{-1}        // 使用栈存储状态（起始索引）
	path := make([]int, 0, k) // 复用路径切片

	return func() []int {
		for len(stack) > 0 {
			// 弹出栈顶状态
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 回溯：移除路径中超过当前状态长度的元素
			if start >= 0 {
				path = path[:len(path)-1]
			}

			// 尝试下一个起始位置
			for i := start + 1; i < n; i++ {
				// 添加新元素到路径
				path = append(path, cand[i])

				// 将当前状态（起始索引）入栈
				stack = append(stack, i)

				// 如果路径长度达到上限，不再扩展
				if len(path) == k {
					break
				}
			}

			// 返回有效路径（非空）
			if len(path) > 0 {
				return path
			}
		}
		return nil // 无更多子集
	}
}
