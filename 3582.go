package main

import "strings"

func generateTag(caption string) string {
	var b []byte
	b = append(b, '#')
	strs := strings.Fields(caption)
	for i := 0; i < len(strs); i++ {
		if i == 0 {
			b = append(b, []byte(strings.ToLower(strs[i]))...)
			continue
		}
		s := strings.ToLower(strs[i])
		for j, c := range s {
			if j == 0 && c >= 'a' && c <= 'z' {
				c += 'A' - 'a'
			}
			b = append(b, byte(c))
		}
	}
	if len(b) > 100 {
		b = b[:100]
	}
	return string(b)
}
