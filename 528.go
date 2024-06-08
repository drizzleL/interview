package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	data []int
	max  int
}

func ConstructorRand(w []int) Solution {
	sol := Solution{}
	for _, v := range w {
		sol.max += v
		sol.data = append(sol.data, sol.max-1)
	}
	return sol
}

func (this *Solution) PickIndex() int {
	v := rand.Intn(this.max)
	return sort.SearchInts(this.data, v)
}
