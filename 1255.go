package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterDict := [26]int{}
	for _, l := range letters {
		letterDict[l-'a'] += 1
	}
	var ret int
	var helper func(i int, point int)
	helper = func(i int, point int) {
		if i == len(words) {
			ret = max(ret, point)
			return
		}
		var flag bool
		for _, c := range words[i] {
			letterDict[c-'a'] -= 1
			if letterDict[c-'a'] < 0 {
				flag = true
			}
			point += score[c-'a']
		}
		if !flag {
			helper(i+1, point)
		}
		for _, c := range words[i] {
			letterDict[c-'a'] += 1
			point -= score[c-'a']
		}
		helper(i+1, point)
	}
	helper(0, 0)
	return ret
}
