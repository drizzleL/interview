package main

func kthDistinct(arr []string, k int) string {
	dict := map[string]int{}
	for _, str := range arr {
		dict[str] += 1
	}
	for _, str := range arr {
		if dict[str] != 1 {
			continue
		}
		k -= 1
		if k == 0 {
			return str
		}
	}
	return ""
}
