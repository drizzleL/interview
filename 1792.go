package main

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	type class struct {
		total int
		pass  int
		add   float64
	}
	getAdd := func(total, pass int) float64 {
		a, b := float64(total), float64(pass)
		return (a - b) / (a * (a + 1))
	}
	q := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(*class).add > b.(*class).add
		},
	}
	for _, cla := range classes {
		heap.Push(&q, &class{
			total: cla[1],
			pass:  cla[0],
			add:   getAdd(cla[1], cla[0]),
		})
	}
	for i := 0; i < extraStudents; i++ {
		top := heap.Pop(&q).(*class)
		top.total += 1
		top.pass += 1
		top.add = getAdd(top.total, top.pass)
		heap.Push(&q, top)
	}
	var ret float64
	getPass := func(total, pass int) float64 {
		a, b := float64(total), float64(pass)
		return b / a
	}
	for q.Len() != 0 {
		top := heap.Pop(&q).(*class)
		ret += getPass(top.total, top.pass)
	}
	return ret / float64(len(classes))
}
