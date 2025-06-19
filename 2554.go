package main

func maxCount(banned []int, n int, maxSum int) int {
	dict := map[int]bool{}
	for _, ban := range banned {
		dict[ban] = true
	}
	var ret, sum int
	for i := 1; i <= n; i++ {
		if dict[i] {
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		ret += 1
	}
	return ret
}
