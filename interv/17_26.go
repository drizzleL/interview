package main

import (
	"fmt"
	"math"
)

func computeSimilarities(docs [][]int) []string {
	size := len(docs)
	sames := make([]int, size*size)
	dict := map[int][]int{}
	for i, doc := range docs {
		for _, v := range doc {
			for _, oldI := range dict[v] {
				sames[oldI*size+i]++
			}
			dict[v] = append(dict[v], i)
		}
	}
	var ret []string
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			same := sames[i*size+j]
			if same == 0 {
				continue
			}
			simi := float64(same) / float64(len(docs[i])+len(docs[j])-same)
			simi = math.Round(simi*1e4) / 1e4
			ret = append(ret, fmt.Sprintf("%d,%d: %.4f", i, j, simi))
		}
	}
	return ret
}
