package main

func numOfUnplacedFruits2(fruits []int, baskets []int) int {
	ret := len(fruits)
	for _, f := range fruits {
		for i, b := range baskets {
			if b < f {
				continue
			}
			baskets[i] = -1
			ret -= 1
			break
		}
	}
	return ret
}
