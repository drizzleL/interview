package main

func dividePlayers(skill []int) int64 {
	var sum int
	for _, sk := range skill {
		sum += sk
	}
	each := sum / (len(skill) / 2)
	if sum != each*(len(skill)/2) {
		return -1
	}
	var ret int
	dict := make([]int, each)
	for _, sk := range skill {
		if sk >= each {
			return -1
		}
		if dict[each-sk] != 0 {
			dict[each-sk] -= 1
			ret += sk * (each - sk)
			continue
		}
		dict[sk] += 1
	}
	for _, v := range dict {
		if v != 0 {
			return -1
		}
	}
	return int64(ret)
}
