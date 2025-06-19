package main

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	i, j := 0, len(plants)-1
	var ret int
	a, b := capacityA, capacityB
	for ; i < j; i, j = i+1, j-1 {
		if a < plants[i] {
			a = capacityA
			ret += 1
		}
		a -= plants[i]
		if b < plants[j] {
			b = capacityB
			ret += 1
		}
		b -= plants[j]
	}
	if i == j && a < plants[i] && b < plants[j] {
		ret += 1
	}
	return ret
}
