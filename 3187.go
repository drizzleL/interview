package main

func countOfPeaks(nums []int, queries [][]int) []int {
	isPeak := func(i int) int {
		if i == 0 || i == len(nums)-1 {
			return 0
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return 1
		}
		return 0
	}
	f := NewFenwick(len(nums))
	queryRange := func(l, r int) int {
		if l > r {
			return 0
		}
		return f.sum(r) - f.sum(l-1)
	}
	update := func(idx int, delta int) {
		f.add(idx, delta)
	}
	p := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		p[i] = isPeak(i)
		update(i, p[i])
	}

	ret := make([]int, 0, len(queries))
	for _, q := range queries {
		if q[0] == 2 {
			idx, val := q[1], q[2]
			nums[idx] = val
			for j := idx - 1; j < idx+2; j++ {
				if j < 0 || j >= len(nums) {
					continue
				}
				old, now := p[j], isPeak(j)
				p[j] = now
				update(j, now-old)
			}
			continue
		}
		l, r := q[1], q[2]
		ret = append(ret, queryRange(l+1, r-1))
	}
	return ret
}
