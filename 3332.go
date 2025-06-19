package main

func maxScore2(n int, k int, stayScore [][]int, travelScore [][]int) int {
	ret := make([][]int, k)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		ret[0][i] = stayScore[0][i]
		for j := 0; j < n; j++ {
			ret[0][i] = max(ret[0][i], travelScore[j][i])
		}
	}
	for day := 1; day < k; day++ {
		for i := 0; i < n; i++ {
			ret[day][i] = ret[day-1][i] + stayScore[day][i]
			for j := 0; j < n; j++ {
				ret[day][i] = max(ret[day][i], ret[day-1][j]+travelScore[j][i])
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, ret[k-1][i])
	}
	return ans
}
