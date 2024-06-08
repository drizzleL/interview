package main

import (
	"math"
)

func distributeCandies(candies int, num_people int) []int {
	round := (int(math.Sqrt(float64(8*candies+1))) - 1) / 2 / num_people
	ret := make([]int, num_people)
	for i := 0; i < num_people; i++ {
		last := i + 1 + (round-1)*num_people
		ret[i] = (i + 1 + last) * round / 2
		candies -= ret[i]
	}
	for i := 0; i < num_people; i++ {
		candy := i + 1 + round*num_people
		candy = min(candy, candies)
		candies -= candy
		ret[i] += candy
		if candies == 0 {
			break
		}
	}
	return ret
}
