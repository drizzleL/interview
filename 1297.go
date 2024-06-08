package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	var ret int
	dict := map[string]int{}
	cnt := make([]int, 26)
	var letterSize int
	for i := 0; i < minSize; i++ {
		cnt[s[i]-'a'] += 1
		if cnt[s[i]-'a'] == 1 {
			letterSize += 1
		}
	}
	if letterSize <= maxLetters {
		dict[s[:minSize]] += 1
	}

	for i := minSize; i < len(s); i++ {
		cnt[s[i]-'a'] += 1
		if cnt[s[i]-'a'] == 1 {
			letterSize += 1
		}
		cnt[s[i-minSize]-'a'] -= 1
		if cnt[s[i-minSize]-'a'] == 0 {
			letterSize -= 1
		}
		if letterSize <= maxLetters {
			dict[s[i-minSize+1:i+1]] += 1
		}
	}
	for _, c := range dict {
		ret = max(ret, c)
	}
	return ret
}
