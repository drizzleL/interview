package main

func maxScoreSightseeingPair(values []int) int {
	tmp := values[0]
	var ret int
	for i := 1; i < len(values); i++ {
		ret = max(ret, values[i]-i+tmp)
		tmp = max(tmp, values[i]+i)
	}
	return ret
}
