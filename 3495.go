package main

func minOperations193(queries [][]int) int64 {
	getPresum := func(maxLimit int) int {
		if maxLimit == 0 {
			return 0
		}
		var ret int
		l, r := 1, 3
		for k := 1; l <= maxLimit; k++ {
			ret += k * (min(maxLimit, r) - l + 1)
			l, r = r+1, (r+1)*4-1
		}
		return ret
	}
	var ret int
	for _, q := range queries {
		l, r := q[0], q[1]
		sum := getPresum(r) - getPresum(l-1)
		ret += (sum + 1) / 2
	}
	return int64(ret)
}
