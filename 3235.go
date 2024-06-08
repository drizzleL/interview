package main

import (
	"log"
	"math"
)

// TODO
func canReachCorner(x, y int, circles [][]int) bool {
	isIn := func(x1, y1, x2, y2 int, r int) bool {
		return (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= r*r
	}
	for _, cir := range circles {
		if isIn(0, 0, cir[0], cir[1], cir[2]) || isIn(x, y, cir[0], cir[1], cir[2]) {
			return false
		}
	}
	parent := make([]int, len(circles))
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	union := func(i, j int) {
		pa, pb := find(i), find(j)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	for i := 0; i < len(circles); i++ {
		for j := i + 1; j < len(circles); j++ {
			if isIn(circles[i][0], circles[i][1], circles[j][0], circles[j][1], circles[i][2]+circles[j][2]) {
				log.Println(circles[i][0]-circles[j][0], circles[i][1]-circles[j][1])
				union(i, j)
			}
		}
	}
	groupCircle := map[int][][]int{}
	for c, p := range parent {
		groupCircle[p] = append(groupCircle[p], circles[c])
	}
	log.Println(groupCircle)
	for _, cirs := range groupCircle {
		left, right, top, down := math.MaxInt32, math.MinInt32, math.MinInt32, math.MaxInt32
		for _, cir := range cirs {
			log.Println(cir)
			if cir[0]-x > cir[2] || 0-cir[0] > cir[2] || cir[1]-y > cir[2] || 0-cir[1] > cir[2] {
				continue
			}
			if cir[0] > x && cir[1] > y && !isIn(cir[0], cir[1], x, y, cir[2]) {
				continue
			}
			if cir[0] < 0 && cir[1] < 0 && !isIn(cir[0], cir[1], 0, 0, cir[2]) {
				continue
			}
			if cir[0] < 0 && cir[1] > y && !isIn(cir[0], cir[1], 0, y, cir[2]) {
				continue
			}
			if cir[0] > x && cir[1] < 0 && !isIn(cir[0], cir[1], x, 0, cir[2]) {
				continue
			}
			left = min(left, cir[0]-cir[2])
			right = max(right, cir[0]+cir[2])
			down = min(down, cir[1]-cir[2])
			top = max(top, cir[1]+cir[2])
		}
		if left <= 0 && right >= x || top >= y && down <= 0 {
			return false
		}
		if left <= 0 && down <= 0 || top >= y && right >= x {
			return false
		}
	}
	return true
}
