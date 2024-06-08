package main

func maxPoints(points [][]int) int64 {
	if len(points) == 0 {
		return 0
	}
	last := points[0]
	sum := make([]int, len(points[0]))
	for i := 1; i < len(points); i++ {
		for j, p := range points[i] {
			sum[j] = last[j] + p
		}
		maxLeft := last[0] - 1
		for j := 1; j < len(points[0]); j++ {
			sum[j] = max(sum[j], points[i][j]+maxLeft)
			maxLeft = max(maxLeft, last[j]) - 1
		}
		maxRight := last[len(last)-1] - 1
		for j := len(points[0]) - 1; j >= 0; j-- {
			sum[j] = max(sum[j], points[i][j]+maxRight)
			maxRight = max(maxRight, last[j]) - 1
		}
		last, sum = sum, last
	}
	var ret int
	for _, c := range last {
		ret = max(ret, c)
	}
	return int64(ret)
}
