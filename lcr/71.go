package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	sum  int
	data []int
}

func SolConstructor(w []int) Solution {
	data := []int{}
	var sum int
	for _, v := range w {
		sum += v
		data = append(data, sum)
	}
	return Solution{
		sum:  sum,
		data: data,
	}
}

func (this *Solution) PickIndex() int {
	v := rand.Intn(this.sum)
	return sort.Search(len(this.data), func(i int) bool {
		return this.data[i] > v
	})
}
