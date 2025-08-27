package main

import (
	"log"
	"sort"
)

func maxPartitionsAfterOperations(s string, k int) int {
	if k == 26 {
		return 1
	}
	var seen [26]bool
	var cnt int
	var idxs [26][]int
	suffix := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1]
		c := int(s[i] - 'a')
		idxs[c] = append(idxs[c], i)
		if seen[c] {
			continue
		}
		log.Println(i, c, cnt)
		if cnt%k == 0 {
			suffix[i] += 1
			seen = [26]bool{}
			cnt = 0
		}
		cnt += 1
		seen[c] = true
	}
	for i := range idxs {
		sort.Ints(idxs[i])
	}
	var precnt int
	cnt = 0
	ret := suffix[0]
	seen = [26]bool{}
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		idxs[c] = idxs[c][1:]
		for j := 0; j < 26; j++ {
			if j == c || seen[j] {
				continue
			}
			tmp := precnt
			var nextIdxs []int
			var nextIdx int
			if cnt == k { // start new session
				tmp += 1
				nextIdx = k - 1
				for j2 := 0; j2 < 26; j2++ {
					if j2 == j || len(idxs[j2]) == 0 {
						continue
					}
					nextIdxs = append(nextIdxs, idxs[j2][0])
				}
			} else {
				nextIdx = k - cnt - 1
				for j2 := 0; j2 < 26; j2++ {
					if j2 == j || seen[j2] || len(idxs[j2]) == 0 {
						continue
					}
					nextIdxs = append(nextIdxs, idxs[j2][0])
				}
			}
			if nextIdx < len(nextIdxs) {
				sort.Ints(nextIdxs)
				tmp += suffix[nextIdxs[nextIdx]]
			}
			ret = max(ret, 1+tmp)
		}
		if seen[c] {
			continue
		}
		if i != 0 && cnt%k == 0 {
			precnt += 1
			seen = [26]bool{}
			cnt = 0
		}
		cnt += 1
		seen[c] = true
	}
	return ret
}
