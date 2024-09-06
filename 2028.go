package main

func missingRolls(rolls []int, mean int, n int) []int {
	var sum int
	for _, ro := range rolls {
		sum += ro
	}
	sum = mean*(len(rolls)+n) - sum
	if sum < n || sum > n*6 {
		return nil
	}
	ret := make([]int, n)
	for i := range ret {
		ret[i] = 1
	}
	sum -= len(ret)
	for i := 0; i < len(ret) && sum > 0; i++ {
		add := min(5, sum)
		ret[i] += add
		sum -= add
	}
	return ret
}
