package main

import "container/heap"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var h PrimeEles
	for i := 0; i < len(arr)-1; i++ {
		heap.Push(&h, []int{i, len(arr) - 1, arr[i], arr[len(arr)-1]})
	}
	for i := 1; i < k; i++ {
		top := heap.Pop(&h).([]int)
		top[1]--
		if top[1] == top[0] {
			continue
		}
		top[3] = arr[top[1]]
		heap.Push(&h, top)
	}
	v := heap.Pop(&h).([]int)
	return v[2:]
}

type PrimeEles [][]int

func (pe PrimeEles) Len() int {
	return len(pe)
}

func (pe *PrimeEles) Pop() any {
	v := (*pe)[pe.Len()-1]
	*pe = (*pe)[:pe.Len()-1]
	return v
}

func (pe *PrimeEles) Push(x any) {
	*pe = append(*pe, x.([]int))
}

func (pe PrimeEles) Less(i, j int) bool {
	a1, b1 := pe[i][2], pe[i][3]
	a2, b2 := pe[j][2], pe[j][3]
	return float64(a1)/float64(b1) < float64(a2)/float64(b2)
}

func (pe PrimeEles) Swap(i, j int) {
	pe[i], pe[j] = pe[j], pe[i]
}
