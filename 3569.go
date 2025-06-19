package main

func maximumCount2(nums []int, queries [][]int) []int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]bool, maxVal+1)
	sieve[1] = true
	for i := 2; i <= maxVal; i++ {
		if sieve[i] {
			continue
		}
		for j := i * 2; j <= maxVal; j += i {
			sieve[j] = true
		}
	}
	ret := make([]int, len(queries))
	return ret
}
