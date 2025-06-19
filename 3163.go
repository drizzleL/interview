package main

func compressedString(word string) string {
	var cnt int
	var b []byte
	for i, c := range word {
		cnt += 1
		if i == len(word)-1 || word[i+1] != byte(c) || cnt == 9 {
			b = append(b, '0'+byte(cnt))
			b = append(b, byte(c))
			cnt = 0
		}
	}
	return string(b)
}
