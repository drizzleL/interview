package main

import "container/heap"

func minimizeStringValue(s string) string {
	dict := make([]int, 26)
	var k int
	for _, c := range s {
		if c == '?' {
			k += 1
			continue
		}
		dict[c-'a'] += 1
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			av, bv := a.([2]int), b.([2]int)
			if av[1] == bv[1] {
				return av[0] < bv[0]
			}
			return av[1] < bv[1]
		},
	}
	for i, cnt := range dict {
		heap.Push(h, [2]int{i, cnt})
	}
	cnts := make([]int, 26)
	for k != 0 {
		v := heap.Pop(h).([2]int)
		cnts[v[0]] += 1
		v[1] += 1
		heap.Push(h, v)
		k -= 1
	}
	ret := make([]byte, len(s))
	var idx int
	for i := range s {
		if s[i] != '?' {
			ret[i] = s[i]
			continue
		}
		for idx < len(cnts) && cnts[idx] == 0 {
			idx += 1
		}
		ret[i] = 'a' + byte(idx)
		cnts[idx] -= 1
	}
	return string(ret)
}
