package main

func clearStars(s string) string {
	var dict [26][]int
	hide := make([]bool, len(s))
	for i, c := range s {
		if c != '*' {
			dict[c-'a'] = append(dict[c-'a'], i)
			continue
		}
		hide[i] = true
		for j := range dict {
			if len(dict[j]) == 0 {
				continue
			}
			last := dict[j][len(dict[j])-1]
			dict[j] = dict[j][:len(dict[j])-1]
			hide[last] = true
			break
		}
	}
	var b []byte
	for i, c := range s {
		if hide[i] {
			continue
		}
		b = append(b, byte(c))
	}
	return string(b)
}
