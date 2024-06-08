package main

func groupAnagrams(strs []string) [][]string {
	dict := map[[26]int]int{}
	var ret [][]string
	for _, s := range strs {
		v := [26]int{}
		for _, c := range s {
			v[c-'a'] += 1
		}
		if _, exist := dict[v]; !exist {
			dict[v] = len(ret)
			ret = append(ret, nil)
		}
		ret[dict[v]] = append(ret[dict[v]], s)
	}
	return ret
}
