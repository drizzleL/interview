package main

func minOperations13(queries [][]int) int64 {
	getPresum := func(x int) int {
		if x == 0 {
			return 0
		}
		l, r := 1, 3
		var ret int
		for k := 1; l <= x; k++ {
			ret += (min(x, r) - l + 1) * k
			l, r = r+1, (r+1)*4-1
		}
		return ret
	}
	getSum := func(l, r int) int {
		return getPresum(r) - getPresum(l-1)
	}
	var ret int
	for _, q := range queries {
		l, r := q[0], q[1]
		sum := getSum(l, r)
		ret += sum / 2
		if sum/2*2 != sum {
			ret += 1
		}
	}
	return int64(ret)
}
