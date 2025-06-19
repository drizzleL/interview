package main

import "sort"

func topKFrequent(words []string, k int) []string {
	sort.Strings(words)
	type word struct {
		word string
		freq int
	}
	var ww []*word
	for i, w := range words {
		if i != 0 && w == words[i-1] {
			ww[len(ww)-1].freq += 1
			continue
		}
		ww = append(ww, &word{
			word: w,
			freq: 1,
		})
	}
	sort.Slice(ww, func(i, j int) bool {
		if ww[i].freq == ww[j].freq {
			return ww[i].word < ww[j].word
		}
		return ww[i].freq > ww[j].freq
	})
	var ret []string
	for i := 0; i < k; i++ {
		ret = append(ret, ww[i].word)
	}
	return ret

}
