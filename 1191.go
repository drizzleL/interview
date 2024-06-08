package main

func kConcatenationMaxSum(arr []int, k int) int {
	var midsum int
	var sum int
	var tmp int
	for _, v := range arr {
		sum += v
		tmp = max(0, tmp+v)
		midsum = max(midsum, tmp)
	}
	if k <= 1 {
		return midsum * k % (1e9 + 7)
	}

	var presum int
	tmp = 0
	for i := 0; i < len(arr); i++ {
		tmp += arr[i]
		presum = max(presum, tmp)
	}
	var sufsum int
	tmp = 0
	for i := len(arr) - 1; i >= 0; i-- {
		tmp += arr[i]
		sufsum = max(sufsum, tmp)
	}
	ret := max(sufsum+presum+max(0, sum*(k-2)), midsum)
	ret %= 1e9 + 7
	return ret
}
