package main

import "math"

func longestBalanced(nums []int) int {
	helper := func(v int) int { // odd:1 even: -1
		return (v%2)*2 - 1
	}
	next := make([]int, len(nums))
	for i := range next {
		next[i] = len(nums)
	}
	seen := map[int]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if j, ok := seen[nums[i]]; ok {
			next[i] = j
		}
		seen[nums[i]] = i
	}
	seen = map[int]int{}
	score := make([]int, len(nums))
	for i, num := range nums {
		if i != 0 {
			score[i] = score[i-1]
		}
		if _, ok := seen[num]; !ok {
			score[i] += helper(num)
		}
		seen[num] = i
	}
	segTree := make([][3]int, len(nums)*4) // 0:min 1:max 2:lazy
	for i := range segTree {
		segTree[i][0] = math.MaxInt32
		segTree[i][1] = math.MinInt32
	}
	var build func(i, l, r int)
	build = func(i, l, r int) {
		if l == r {
			segTree[i][0] = score[l]
			segTree[i][1] = score[l]
			return
		}
		mid := l + (r-l)/2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		segTree[i][0] = min(segTree[i*2][0], segTree[i*2+1][0])
		segTree[i][1] = max(segTree[i*2][1], segTree[i*2+1][1])
	}
	build(1, 0, len(score)-1)
	pushDown := func(i, l, r int) {
		if segTree[i][2] == 0 {
			return
		}
		segTree[i][0] += segTree[i][2]
		segTree[i][1] += segTree[i][2]
		if l != r {
			segTree[i*2][2] += segTree[i][2]
			segTree[i*2+1][2] += segTree[i][2]
		}
		segTree[i][2] = 0
	}
	var findRightMostZero func(i, l, r int) int
	findRightMostZero = func(i, l, r int) int {
		pushDown(i, l, r)
		if segTree[i][0] > 0 || segTree[i][1] < 0 {
			return -1
		}
		if l == r {
			return l
		}
		mid := l + (r-l)/2
		if j := findRightMostZero(i*2+1, mid+1, r); j != -1 {
			return j
		}
		return findRightMostZero(i*2, l, mid)
	}
	var update func(start, end int, i, l, r, val int)
	update = func(start, end int, i, l, r, val int) {
		pushDown(i, l, r)
		if l > end || r < start || l > r {
			return
		}
		if l >= start && r <= end {
			segTree[i][2] += val
			pushDown(i, l, r)
			return
		}
		mid := l + (r-l)/2
		update(start, end, i*2, l, mid, val)
		update(start, end, i*2+1, mid+1, r, val)
		segTree[i][0] = min(segTree[i*2][0], segTree[i*2+1][0])
		segTree[i][1] = max(segTree[i*2][1], segTree[i*2+1][1])
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		j := findRightMostZero(1, 0, len(nums)-1)
		if j != -1 {
			ret = max(ret, j-i+1)
		}
		nxt := next[i]
		update(i, nxt-1, 1, 0, len(nums)-1, -helper(nums[i]))
	}
	return ret
}

func longestBalanced2(nums []int) int {
	helper := func(v int) int { // odd:1 even: -1
		return (v%2)*2 - 1
	}
	segTree := make([][3]int, len(nums)*4) // 0:min 1:max 2:lazy
	pushDown := func(i, l, r int) {
		if segTree[i][2] == 0 {
			return
		}
		segTree[i][0] += segTree[i][2]
		segTree[i][1] += segTree[i][2]
		if l != r {
			segTree[i*2][2] += segTree[i][2]
			segTree[i*2+1][2] += segTree[i][2]
		}
		segTree[i][2] = 0
	}
	var findLeftMostZero func(i, l, r int) int
	findLeftMostZero = func(i, l, r int) int {
		pushDown(i, l, r)
		if segTree[i][0] > 0 || segTree[i][1] < 0 {
			return -1
		}
		if l == r {
			return l
		}
		mid := l + (r-l)/2
		if j := findLeftMostZero(i*2, l, mid); j != -1 {
			return j
		}
		return findLeftMostZero(i*2, mid+1, r)
	}
	var update func(start, end int, i, l, r, val int)
	update = func(start, end int, i, l, r, val int) {
		pushDown(i, l, r)
		if l > end || r < start || l > r {
			return
		}
		if l >= start && r <= end {
			segTree[i][2] += val
			pushDown(i, l, r)
			return
		}
		mid := l + (r-l)/2
		update(start, end, i*2, l, mid, val)
		update(start, end, i*2+1, mid+1, r, val)
		segTree[i][0] = min(segTree[i*2][0], segTree[i*2+1][0])
		segTree[i][1] = max(segTree[i*2][1], segTree[i*2+1][1])
	}
	var ret int
	last := map[int]int{}
	getLast := func(num int) int {
		v, ok := last[num]
		if ok {
			return v
		}
		return -1
	}
	for j := 0; j < len(nums); j++ {
		update(getLast(nums[j])+1, j, 1, 0, len(nums)-1, helper(nums[j]))
		i := findLeftMostZero(1, 0, len(nums)-1)
		if i != -1 {
			ret = max(ret, j-i+1)
		}
		last[nums[j]] = j
	}
	return ret
}
