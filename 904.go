package main

func totalFruit(fruits []int) int {
	var ret int
	h := map[int]int{}
	for i, j := 0, 0; j < len(fruits); i++ {
		h[fruits[j]] += 1
		for len(h) > 2 {
			h[fruits[i]] -= 1
			if h[fruits[i]] == 0 {
				delete(h, fruits[i])
			}
			i += 1
		}
		ret = max(ret, j-i+1)
	}
	return ret
}
