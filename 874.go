package main

import (
	"sort"
)

func robotSim(commands []int, obstacles [][]int) int {
	obsCol := map[int][]int{}
	obsRow := map[int][]int{}
	for _, obs := range obstacles {
		obsCol[obs[0]] = append(obsCol[obs[0]], obs[1])
		obsRow[obs[1]] = append(obsRow[obs[1]], obs[0])
	}
	for i := range obsRow {
		sort.Ints(obsRow[i])
	}
	for i := range obsCol {
		sort.Ints(obsCol[i])
	}
	check := func(obs []int, x1, x2 int) int {
		if len(obs) == 0 || x1 == x2 {
			return x2
		}
		idx := sort.SearchInts(obs, x1)
		if idx != len(obs) && obs[idx] == x1 { // right on obstacle
			if x2 > x1 {
				idx = sort.SearchInts(obs, x1+1)
			} else {
				idx = sort.SearchInts(obs, x1-1)
			}
		}
		if x2 > x1 {
			if idx == len(obs) || obs[idx] > x2 {
				return x2
			}
			return obs[idx] - 1
		}
		if idx == 0 || obs[idx-1] < x2 {
			return x2
		}
		return obs[idx-1] + 1
	}
	x, y := 0, 0
	// up:1,down:2,left:3,right:4
	dir := 1
	dirDict1 := map[int]int{
		1: 3,
		2: 4,
		3: 2,
		4: 1,
	}
	dirDict2 := map[int]int{
		1: 4,
		2: 3,
		3: 1,
		4: 2,
	}
	var ret int
	for _, c := range commands {
		var mv int
		switch c {
		case -2:
			dir = dirDict1[dir]
		case -1:
			dir = dirDict2[dir]
		default:
			mv = c
		}
		switch dir {
		case 1:
			y = check(obsCol[x], y, y+mv)
		case 2:
			y = check(obsCol[x], y, y-mv)
		case 3:
			x = check(obsRow[y], x, x-mv)
		case 4:
			x = check(obsRow[y], x, x+mv)
		}
		ret = max(ret, x*x+y*y)
	}
	return ret
}
