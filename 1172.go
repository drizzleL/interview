package main

import (
	"container/heap"
)

type DinnerPlates struct {
	plates   [][]int
	capacity int
	h        PriorityQueue
}

func DinnerConstructor(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
		h:        PriorityQueue{},
	}
}

func (this *DinnerPlates) Push(val int) {
	if this.h.Len() == 0 {
		heap.Push(&this.h, &Item{
			idx:  len(this.plates),
			cost: len(this.plates),
		})
		this.plates = append(this.plates, nil)
	}
	idx := this.h[0].idx
	if len(this.plates) == idx {
		this.plates = append(this.plates, nil)
	}
	this.plates[idx] = append(this.plates[idx], val)
	if len(this.plates[idx]) == this.capacity {
		heap.Pop(&this.h)
	}
}

func (this *DinnerPlates) Pop() int {
	for len(this.plates) != 0 && len(this.plates[len(this.plates)-1]) == 0 {
		this.plates = this.plates[:len(this.plates)-1]
	}
	if len(this.plates) == 0 {
		return -1
	}
	i := len(this.plates) - 1
	if len(this.plates[i]) == this.capacity {
		heap.Push(&this.h, &Item{
			idx:  i,
			cost: i,
		})
	}
	val := this.plates[i][len(this.plates[i])-1]
	this.plates[i] = this.plates[i][:len(this.plates[i])-1]
	return val
}

func (this *DinnerPlates) PopAtStack(i int) int {
	if i >= len(this.plates) {
		return -1
	}
	if len(this.plates[i]) == 0 {
		return -1
	}
	if len(this.plates[i]) == this.capacity {
		heap.Push(&this.h, &Item{
			idx:  i,
			cost: i,
		})
	}
	val := this.plates[i][len(this.plates[i])-1]
	this.plates[i] = this.plates[i][:len(this.plates[i])-1]
	return val
}
