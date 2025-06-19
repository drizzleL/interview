package main

func canArrange(arr []int, k int) bool {
	dict := make([]int, k)
	for _, num := range arr {
		v := num % k
		if v < 0 {
			v += k
		}
		dict[v] += 1
		if v == 0 {
			if dict[v] != 0 {
				dict[v] -= 1
			}
		} else if dict[k-v] != 0 {
			dict[k-v] -= 1
			dict[v] -= 1
		}
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}
