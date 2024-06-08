package main

func maxFrequencyElements(nums []int) int {
	freqs := make([]int, 101)
	var maxFreq int
	for _, num := range nums {
		freqs[num]++
		maxFreq = max(maxFreq, freqs[num])
	}
	var ret int
	for _, freq := range freqs {
		if freq == maxFreq {
			ret += maxFreq
		}
	}
	return ret
}
