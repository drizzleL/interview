package main

func maxSumAfterPartitioning(arr []int, k int) int {
	sum := make([]int, len(arr)+1)
	for i := range arr {
		var subMax int
		for size := 1; size <= k && i+size < len(sum); size++ {
			subMax = max(subMax, arr[i+size-1])
			sum[i+size] = max(sum[i+size], sum[i]+subMax*size)
		}
	}
	return sum[len(arr)]
}
