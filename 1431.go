package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxVal int
	for _, cand := range candies {
		maxVal = max(maxVal, cand)
	}
	ret := make([]bool, len(candies))
	for i, cand := range candies {
		ret[i] = cand+extraCandies >= maxVal
	}
	return ret
}
