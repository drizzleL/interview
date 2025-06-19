package main

func distributeCandies2(n int, limit int) int64 {
	start := max(0, n-limit*2)
	end := min(limit, n)
	var ret int
	for i := start; i <= end; i++ {
		left := n - i
		minVal := max(0, left-limit)
		maxVal := min(left, limit)
		ret += maxVal - minVal + 1
	}
	return int64(ret)
}
