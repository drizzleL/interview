package main

func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	dict := map[[26]int][]string{}
	for _, str := range strs {
		m := [26]int{}
		for _, c := range str {
			m[c-'a']++
		}
		dict[m] = append(dict[m], str)
	}
	for _, ss := range dict {
		ret = append(ret, ss)
	}
	return ret
}
