package main

func majorityFrequencyGroup(s string) string {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	var freq, maxCnt int
	var ret []int
	group := map[int][]int{}
	for i, c := range dict {
		if c == 0 {
			continue
		}
		group[c] = append(group[c], i)
	}
	for c, v := range group {
		if len(v) < maxCnt {
			continue
		}
		if len(v) > maxCnt {
			freq = c
			maxCnt = len(v)
			ret = v
			continue
		}
		if c < freq {
			continue
		}
		freq = c
		ret = v
	}
	var b []byte
	for _, i := range ret {
		b = append(b, byte(i+'a'))
	}
	return string(b)
}
