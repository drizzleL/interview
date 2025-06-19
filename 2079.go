package main

func wateringPlants(plants []int, capacity int) int {
	cnt := capacity
	var ret int
	for i, p := range plants {
		ret += 1
		cnt -= p
		if i+1 < len(plants) && plants[i+1] < cnt {
			ret += (i + 1) * 2
			cnt = capacity
		}
	}
	return ret
}
