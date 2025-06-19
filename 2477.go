package main

import "math"

func minimumFuelCost(roads [][]int, seats int) int64 {
	dict := map[int][]int{}
	out := make([]int, len(roads)+1)
	for _, r := range roads {
		dict[r[0]] = append(dict[r[0]], r[1])
		dict[r[1]] = append(dict[r[1]], r[0])
		out[r[0]] += 1
		out[r[1]] += 1
	}
	var ends []int
	in := make([]int, len(roads)+1)
	for i := range in {
		in[i] = 1
	}
	for from, toCnt := range out {
		if from == 0 {
			continue
		}
		if toCnt == 1 {
			ends = append(ends, from)
		}
	}
	seen := make([]bool, len(roads)+1)
	var ret int
	for len(ends) != 0 {
		end := ends[len(ends)-1]
		ends = ends[:len(ends)-1]
		ret += int(math.Ceil(float64(in[end]) / float64(seats)))
		seen[end] = true
		for _, next := range dict[end] {
			if seen[next] {
				continue
			}
			if next == 0 {
				continue
			}
			in[next] += 1
			out[next] -= 1
			if out[next] == 1 {
				ends = append(ends, next)
			}
		}
	}
	return int64(ret)
}
