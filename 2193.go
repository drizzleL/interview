package main

import "log"

func minMovesToMakePalindrome(s string) int {
	dict := map[byte][]int{}
	for i, c := range s {
		dict[byte(c)] = append(dict[byte(c)], i)
	}
	l, r := 0, len(s)-1
	var ret int
	used := make([]bool, len(s))
	remove := func(v byte) {
		size := len(dict[v])
		dict[v] = dict[v][1 : size-1]
	}
	log.Println(s)
	for l < r-1 {
		log.Println(s[l], s[r], l, r, ret)
		if s[l] == s[r] {
			used[l] = true
			used[r] = true
			dict[s[l]] = dict[s[l]][1:]
			dict[s[r]] = dict[s[r]][:len(dict[s[r]])-1]
			for l < r-1 && used[l] {
				l++
			}
			for l < r-1 && used[r] {
				r--
			}
			continue
		}
		if len(dict[s[l]]) >= len(dict[s[r]]) { // use left
			c := s[l]
			used[l] = true
			lastIdx := dict[c][len(dict[c])-1]
			used[lastIdx] = true
			for l < r-1 && used[l] {
				l++
			}
			for ; lastIdx <= r; lastIdx++ {
				if !used[lastIdx] {
					ret += 1
				}
			}
			remove(c)
		} else {
			c := s[r]
			used[r] = true
			first := dict[c][0]
			used[first] = true
			for l < r-1 && used[r] {
				r--
			}
			for ; first >= l; first-- {
				if !used[first] {
					ret += 1
				}
			}
			remove(c)
		}
	}
	return ret
}
