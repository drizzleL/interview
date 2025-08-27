package main

func partitionString2(s string) []string {
	var ret []string
	seen := map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if seen[s[i:j+1]] {
				continue
			}
			ret = append(ret, s[i:j+1])
			seen[s[i:j+1]] = true
			i = j
		}
	}
	return ret
}
