package main

func countPairs6(nums []int, k int) int64 {
	dict := map[int]int{}
	var ret int
	for _, num := range nums {
		dict[gcd(k, num)] += 1
	}
	for k1, v1 := range dict {
		for k2, v2 := range dict {
			if k2 < k1 {
				continue
			}
			if (k1*k2)%k != 0 {
				continue
			}
			if k1 != k2 {
				ret += v1 * v2
			} else {
				ret += v1 * (v2 - 1) / 2
			}
		}
	}
	return int64(ret)
}
