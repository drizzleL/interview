package main

func minArrivalsToDiscard(arrivals []int, w int, m int) int {
	var ret int
	cnt := map[int]int{}
	discard := make([]bool, len(arrivals))
	for i := 0; i < w; i++ {
		item := arrivals[i]
		cnt[item] += 1
		if cnt[item] > m {
			discard[i] = true
			ret += 1
		}
	}
	for i := w; i < len(arrivals); i++ {
		l := arrivals[i-w]
		if !discard[i-w] {
			cnt[l] -= 1
		}
		r := arrivals[i]
		cnt[r] += 1
		if cnt[r] > m {
			discard[i] = true
			ret += 1
		}
	}
	return ret
}
