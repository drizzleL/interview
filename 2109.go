package main

func addSpaces(s string, spaces []int) string {
	var b []byte
	var idx int
	for i, c := range s {
		if idx < len(spaces) && spaces[idx] == i {
			b = append(b, ' ')
			idx += 1
		}
		b = append(b, byte(c))
	}
	return string(b)
}
