package main

func maximumLength5(nums []int, k int) int {
	cnts := make([][]int, k)
	for i := range cnts {
		cnts[i] = make([]int, k)
	}
	var ret int
	for _, num := range nums {
		for i := 0; i < k; i++ {
			j := (num + i) % k
			cnts[j][num%k] = cnts[j][i] + 1
			ret = max(ret, cnts[j][num%k])
		}
	}
	return ret
}
