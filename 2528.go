package main

import (
	"math"
)

func maxPower(stations []int, r int, k int) int64 {
	maxPower, minPower := math.MinInt32, math.MaxInt32
	powers := make([]int, len(stations))
	var lsum, rsum int
	for i := 1; i <= r; i++ {
		rsum += stations[i]
	}
	for i := 0; i < len(stations); i++ {
		powers[i] = lsum + stations[i] + rsum
		minPower = min(minPower, powers[i])
		maxPower = max(maxPower, powers[i])
		lsum += stations[i]
		if i-r >= 0 {
			lsum -= stations[i-r]
		}
		if i+1 < len(stations) {
			rsum -= stations[i+1]
		}
		if i+1+r < len(stations) {
			rsum += stations[i+1+r]
		}
	}
	maxPower += k
	check := func(p int) bool {
		var used int
		var tmpsum int
		rb := NewRb(r * 2)
		for i := 0; i < len(powers); i++ {
			if p-tmpsum-powers[i]+used > k {
				return false
			}
			if tmpsum+powers[i] < p {
				add := p - tmpsum - powers[i]
				tmpsum -= rb.Push(add)
				tmpsum += add
				used += add
			} else {
				tmpsum -= rb.Push(0)
			}
		}
		return true
	}
	for minPower < maxPower {
		mid := (minPower+maxPower)/2 + 1
		if check(mid) {
			minPower = mid
		} else {
			maxPower = mid - 1
		}
	}
	return int64(minPower)
}

func NewRb(size int) *Rb {
	return &Rb{data: make([]int, size)}
}

type Rb struct {
	data []int
	idx  int
}

func (r *Rb) Push(v int) int {
	if len(r.data) == 0 {
		return v
	}
	old := r.data[r.idx]
	r.data[r.idx] = v
	r.idx += 1
	r.idx %= len(r.data)
	return old
}
