package main

import (
	"math"
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	ret := make([]int, len(queries))
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] > rooms[j][1]
	})
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] > queries[j][1]
	})
	tr := redblacktree.NewWithIntComparator()
	for _, q := range queries {
		preId, minSize, qId := q[0], q[1], q[2]
		for len(rooms) > 0 && rooms[0][1] >= minSize {
			tr.Put(rooms[0][0], rooms[0][0])
			rooms = rooms[1:]
		}
		ret[qId] = -1
		if tr.Size() == 0 {
			continue
		}
		diff := math.MaxInt32
		if v, ok := tr.Floor(preId); ok {
			diff = preId - v.Key.(int)
			ret[qId] = v.Key.(int)
		}
		if v, ok := tr.Ceiling(preId); ok && v.Key.(int)-preId < diff {
			ret[qId] = v.Key.(int)
		}
	}
	return ret
}
